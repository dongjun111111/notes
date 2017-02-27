package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	APPID     = ""
	APPSECRET = ""
)

func HttpPost(url string, bytesReq []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesReq))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	bytesResp, _ := ioutil.ReadAll(resp.Body)
	return bytesResp, nil
}

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type WxToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type UserInfo struct {
	OpenId     string `json:"openid"`
	NickName   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	HeadImgURL string `json:"headimgurl"`
	UnionId    string `json:"unionid"`
}

type ReturnMessage struct {
	errcode int
	errmsg  string
}

//获取openid和access_token
func GetOpenIdandAccessToken(code string) (string, string, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?"
	url += "appid=" + APPID
	url += "&secret=" + APPSECRET
	url += "&code=" + code
	url += "&grant_type=" + "authorization_code"

	bytesResp, err := utils.HttpGet(url)
	if err != nil {
		return "", "", err
	}
	var token WxToken
	json.Unmarshal(bytesResp, &token)

	return token.OpenId, token.AccessToken, nil
}

//获取个人信息
func GetWechatUserInfo(openid, accessToken string) (UserInfo, error) {
	url := "https://api.weixin.qq.com/sns/userinfo?"
	url += "access_token=" + accessToken
	url += "&openid=" + openid

	var w UserInfo
	bytesResp, err := utils.HttpGet(url)
	if err != nil {
		return w, err
	}

	err = json.Unmarshal(bytesResp, &w)
	if err != nil {
		return w, err
	}
	return w, nil
}

//校验openid
func CheckOpenId(code, openid string) bool {

	_, oid, _ := GetOpenIdandAccessToken(code)
	if oid == openid {
		return true
	}
	return false
}

//校验access_token
func CheckAccessToken(accessToken, openid string) (bool, error) {
	url := "https://api.weixin.qq.com/sns/auth?"
	url += "access_token=" + accessToken
	url += "&openid=" + openid

	bytesResp, err := utils.HttpGet(url)
	if err != nil {
		return false, err
	}
	var m ReturnMessage
	err = json.Unmarshal(bytesResp, &m)
	if err != nil {
		return false, err
	}
	if m.errcode == 0 {
		return true, nil
	}
	return false, nil
}
