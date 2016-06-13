package main

import (
	"archive/tar"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//aliyun返回数据类型
type BinLogFile struct {
	HostInstanceID       int    `json:"HostInstanceID"`
	FileSize             int    `json:"FileSize"`
	Checksum             string `json:"Checksum"`
	LinkExpiredTime      string `json:"LinkExpiredTime"`
	LogEndTime           string `json:"LogEndTime"`
	LogBeginTime         string `json:"LogBeginTime"`
	IntranetDownloadLink string `json:"IntranetDownloadLink"`
	DownloadLink         string `json:"DownloadLink"`
}

type Items struct {
	BinLogFiles []BinLogFile `json:"BinLogFile"`
}

type Data struct {
	Item             Items  `json:"Items"`
	PageNumber       int    `json:"PageNumber"`
	TotalRecordCount int    `json:"TotalRecordCount"`
	TotalFileSize    string `json:"TotalFileSize"`
	RequestId        string `json:"RequestId"`
	PageRecordCount  int    `json:"PageRecordCount"`
}

func (this *Data) GetDloadAddr() *[]string {
	var s []string
	if len(this.Item.BinLogFiles) > 0 {
		for _, v := range this.Item.BinLogFiles {
			s = append(s, v.DownloadLink)
		}
	}
	return &s
}

//==========主函数===========================================================================

func main() {
	path := "./"
	for {
		if (time.Now().Hour() == 6) || (time.Now().Hour() == 18) {
			go doOneTime(path)
		}
		time.Sleep(time.Hour)
	}
}

//============================================================================================
//完整的一次下载解压过程
func doOneTime(path string) {
	//权限参数
	accessKeyId := "???"
	accessKey := "???"
	DBInstanceId := "???"
	//aliyun线上实例id
	exampleID := "???"
	//==========================================

	now := time.Now()
	var err error
	_, err = os.Stat(now.Format("20060102"))
	if err != nil {
		os.Mkdir(now.Format("20060102"), 0777)
	}

	//建立日志文件
	var record *os.File
	_, err = os.Stat(now.Format("20060102") + "/downlog")
	if err == nil {
		record, err = os.OpenFile(now.Format("20060102")+"/downlog", os.O_APPEND, 0777)
	} else {
		record, err = os.Create(path + "/" + now.Format("20060102") + "/downlog")
	}
	defer record.Close()
	record.Write([]byte(now.Format("20060102") + "\r\n"))

	//发送请求接受返回参数
	body, err := logApi(accessKeyId, accessKey, DBInstanceId)
	if err != nil {
		record.Write([]byte("aliyun发送请求失败:\r\n"))
		record.Write([]byte(fmt.Sprint(err) + "\r\n"))
		return
	}

	//r获取返回参数
	r := &Data{}
	err = json.Unmarshal(*body, r)
	if err != nil {
		record.Write([]byte("解析参数失败:\r\n"))
		record.Write([]byte(fmt.Sprint(err) + "\r\n"))
		return
	}
	//获取下载列表
	dList := r.GetDloadAddr()
	//选择需要的下载地址
	dLoads := &[]string{}
	for _, v := range *dList {
		if v[63:77] == exampleID {
			*dLoads = append(*dLoads, v)
		}
	}

	for _, v := range *dLoads {
		record.Write([]byte(v + "\r\n"))
		record.Write([]byte("\r\n"))
	}
	//开始下载
	names, err, n := downLoad(dLoads, path, record)
	if err != nil {
		record.Write([]byte(fmt.Sprint(err) + "\r\n"))
		if n < 0 || n == len(*dLoads) {
			return
		}
	}
	//下载完毕开始解压
	for k, v := range *names {
		err := unTar(v, path+"/"+now.Format("20060102"), record)
		if err != nil {
			record.Write([]byte("解压错误: "))
			record.Write([]byte(fmt.Sprint(k, ",", v, ":", err)))
			record.Write([]byte("\r\n"))
		}
		os.Remove(v)
	}
	record.Write([]byte("解压完成\r\n"))
	record.Write([]byte("\r\n\r\n"))
}

func ascllFormat(str string) string {
	var i, j int
	str2 := ""
	for k, v := range str {
		switch string(v) {
		case ":":
			j = k
			str2 += str[i:j] + "%3A"
			i = j + 1
		case "%":
			j = k
			str2 += str[i:j] + "%25"
			i = j + 1
		case "&":
			j = k
			str2 += str[i:j] + "%26"
			i = j + 1
		case "=":
			j = k
			str2 += str[i:j] + "%3D"
			i = j + 1
		case "+":
			j = k
			str2 += str[i:j] + "%2B"
			i = j + 1
		case "/":
			j = k
			str2 += str[i:j] + "%2F"
			i = j + 1
		}
	}
	str2 += str[i:]
	return str2
}

//生成随机数19位
func randInt19ToStr() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprint(rand.Int63())
}

//hmac+sha1+key加密
func hmacSha1(stringToSign, key string) (Signature string) {
	key += "&"
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(stringToSign))
	Signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}

func logApi(accessKeyId, accessKey, DBInstanceId string) (*[]byte, error) {
	format := "json" //强制

	start := time.Now().Add(-time.Hour * 12)
	//唯一随机数
	signatureNonce := randInt19ToStr()
	//UTC开始和截止时间
	startTime := start.UTC().Format("2006-01-02T15:04:05Z07:00")
	endTime := time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
	//addr
	urlAddr := "https://rds.aliyuncs.com?"
	//公共和binlog和私有参数格式化
	para := `AccessKeyId=` + accessKeyId + `&Action=DescribeBinlogFiles&DBInstanceId=` + DBInstanceId + `&EndTime=` + ascllFormat(endTime) + "&Format=" + format + `&SignatureMethod=HMAC-SHA1&SignatureNonce=` + signatureNonce + `&SignatureVersion=1.0&StartTime=` + ascllFormat(startTime) + `&Timestamp=` + ascllFormat(time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")) + `&Version=2014-08-15`

	//生成签名
	var Signature string
	stringToSign := "GET&%2F&" + ascllFormat(para)
	Signature = hmacSha1(stringToSign, accessKey)

	//发送get请求
	var bad []byte
	resp, err := http.Get(urlAddr + para + "&Signature=" + ascllFormat(Signature))
	if err != nil {
		return &bad, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &bad, err
	}
	return &body, nil
}

//通过地址列表下载到指定路径
func downLoad(dLoads *[]string, path string, record *os.File) (names *[]string, err error, n int) {
	l := len(*dLoads)
	names = &[]string{}
	if l > 0 {
		now := time.Now()
		dirInfo, err := os.Stat(now.Format("20060102"))
		if err != nil || !dirInfo.IsDir() {
			os.Mkdir(now.Format("20060102"), 0777)
			dirInfo, _ = os.Stat(now.Format("20060102"))
		}
		dir := dirInfo.Name()
		//多通路下载
		ch := make(chan error, l)
		for k, v := range *dLoads {
			go func() {
				i := k
				//单个下载
				name, err := downOne(v, path+"/"+dir, i, record)
				if err != nil {
					ch <- err
				} else {
					ch <- nil
					*names = append(*names, name)
				}
			}()
			time.Sleep(time.Microsecond)
		}
		//循环等待，记录err
		num := 0
		errs := make([]error, l)
		for {
			if num >= l {
				break
			}
			select {
			case v, _ := <-ch:
				errs[num] = v
				num++

			}
		}
		//全部完成，未验证成功
		record.Write([]byte("all download complete\r\n"))
		record.Write([]byte("\r\n"))
		//错误处理
		errStr := ""
		n := 0 //错误个数
		for k, v := range errs {
			if v != nil {
				n++
				errStr = errStr + fmt.Sprint(k) + ":" + fmt.Sprint(v) + ";"
			}
		}
		if errStr != "" {
			return names, errors.New(errStr), n
		}
		//全部下载，全部成功
		return names, nil, n
	} else { //列表为空时
		return names, errors.New("download list nil"), -1
	}
}

//单个下载
func downOne(url, path string, num int, record *os.File) (name string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	name = path + "/" + time.Now().Format("20060102") + "-" + fmt.Sprint(num) + ".tar"
	out, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader([]byte(body)))
	if err != nil {
		return "", err
	}

	record.Write([]byte(fmt.Sprint("下载完成 ", num, " :"+time.Now().Format("20060102")+"-"+fmt.Sprint(num)+".tar\r\n")))
	return name, nil
}

//解压tar
func unTar(path, pathTo string, record *os.File) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// tar read
	tr := tar.NewReader(f)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// 记录文件
		record.Write([]byte(h.Name + "\r\n"))

		fw, err := os.OpenFile(pathTo+"/"+h.Name, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fw.Close()
		// 复制文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}
