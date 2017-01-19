package webservice

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	DEF_ENCODING = "UTF-8"
	JSON         = "json"
	XML          = "xml"
	POST         = "post"
	GET          = "get"
	CA_JAR       = "CA_JAR"
	CA_WS        = "CA_WS"
	RSA          = "RSA"
)

var (
	WEBSERVICE_URL = "自建WebService服务地址" // webservice请求链接
)

/**
 * 签名返回结果信息
 */
type ResultResponse struct {
	Flag   string
	Result string
}

/**
 * 使用Jar包方式进行CA签名
 * @param msg 处理好后的信息
 * @param platformId 平台号
 * @param password 密码
 * @param encoding 编码
 * @return 签名串
 */
func CaSignEpay(msg, platformId, password, encoding string) string {
	if msg == "" || password == "" || platformId == "" || encoding == "" {
		return ""
	}
	encoding = JSON // 使用json编码
	// ============================SIGN 处理开始====================
	pars := make(map[string]interface{})
	pars["pfx"] = platformId + ".pfx"
	pars["password"] = password
	pars["Message"] = msg
	data, _ := json.Marshal(pars)
	dataStr := string(data)
	dataStr = beego.Htmlquote(dataStr)
	postStr := CreateSOAPXml("http://itrus.com/itrusUtil", "signature", dataStr)
	output := PostWebService(WEBSERVICE_URL, "http://itrus.com/itrusUtil", postStr)
	//===============解析webservice中的xml文件信息===================
	signstr := ""
	var t xml.Token
	var err error
	outputReader := strings.NewReader(output)
	decoder := xml.NewDecoder(outputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.CharData:
			content := string([]byte(token))
			signstr = content
		default:
			// ...
		}
	}
	var ret ResultResponse
	json.Unmarshal([]byte(signstr), &ret)
	if ret.Flag == "true" {
		signstr = ret.Result
	}
	return signstr
}

/**
 * 签名并生成请求URL
 * @param params 参数，签名源串
 * @param platformId 平台号
 * @param password CA证书密钥，RSA签名不需要
 * @param msgType 消息类型，支持json和xml
 * @param apiUrl 请求接口地址
 * @param isMd5 签名串是否为源串的MD5值
 * @param signType 签名类型，CA_JAR、CA_WS
 * @return 签名后的请求地址
 */
func GetReqUrl(params map[string]interface{}, platformId, password, msgType, apiUrl string, isMd5 bool, signType string) (string, error) {
	if params == nil {
		return "", errors.New("请求数据有误~")
	}
	msgType = "json" // 设置默认msg类型
	md5Src := ""     // md5字符串
	sign := ""       // 签名串
	//===============================
	mdata, _ := json.Marshal(params)
	msg := base64.StdEncoding.EncodeToString(mdata)
	if isMd5 {
		md5Src = strings.ToUpper(Md5Encrypt(msg)) // 获取源串的MD5值
	}

	if signType == CA_JAR {
	} else if signType == CA_WS {
		if isMd5 {
			sign = CaSignEpay(md5Src, platformId, password, msgType)
		} else {
			sign = CaSignEpay(msg, platformId, password, msgType)
		}
	}
	if sign != "" {
		ope_url := apiUrl + "?msg=" + msg + "&sign=" + sign + "&msgType=" + msgType
		return ope_url, nil
	} else {
		return "", errors.New("sign处理失败~")
	}
}

/**
 * 使用webservice方式进行验签
 * @param msg 处理好后的信息
 * @param encoding 编码
 * @return 签名串
 */
func VerifyCaSignEpay(verify_msg, verify_sign string) bool {
	if verify_sign == "" {
		return false
	}

	pars := make(map[string]interface{})
	pars["Message"] = verify_msg
	pars["signMessage"] = verify_sign
	data, _ := json.Marshal(pars)
	dataStr := string(data)
	dataStr = beego.Htmlquote(dataStr)
	postStr := CreateSOAPXml("http://itrus.com/itrusUtil", "verify", dataStr)
	output := PostWebService("自建WebService服务地址", "http://itrus.com/itrusUtil", postStr)

	verifyresult := ""
	var t xml.Token
	var err error
	outputReader := strings.NewReader(output)
	decoder := xml.NewDecoder(outputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.CharData:
			content := string([]byte(token))
			verifyresult = content
		default:
			// ...
		}
	}
	var ret ResultResponse
	json.Unmarshal([]byte(verifyresult), &ret)
	if ret.Flag == "true" {
		return true
	}
	return false
}

//POST到webService
func PostWebService(url string, method string, value string) string {
	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	//这里随便传递了点东西
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	//取出主体的内容
	if err != nil {
		fmt.Println("read error", err)
	}
	res.Body.Close()
	return ByteToString(data)
}
func ByteToString(res []byte) string {
	return string(res)
}
func CreateSOAPXml(nameSpace string, methodName string, valueStr string) string {
	soapBody := `<?xml version="1.0" ?><S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">`
	soapBody += "<S:Body>"
	soapBody += "<" + methodName + " xmlns=\"" + nameSpace + "\">"
	soapBody += "<in0>" + valueStr + "</in0>"
	soapBody += "</" + methodName + "></S:Body></S:Envelope>"
	return soapBody
}
