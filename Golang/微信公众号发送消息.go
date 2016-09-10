package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	//发送消息使用导的url
	sendurl = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	//获取token使用导的url
	get_token = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

var requestError = errors.New("request error,check url or network")

type access_token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

//定义一个简单的文本消息格式
type send_msg struct {
	Touser  string            `json:"touser"`
	Toparty string            `json:"toparty"`
	Totag   string            `json:"totag"`
	Msgtype string            `json:"msgtype"`
	Agentid int               `json:"agentid"`
	Text    map[string]string `json:"text"`
	Safe    int               `json:"safe"`
}

type send_msg_error struct {
	Errcode int    `json:"errcode`
	Errmsg  string `json:"errmsg"`
}

func main() {
	mfile := flag.String("m", "", "-m msg.txt 从配置文件读取配置发送消息")
	touser := flag.String("t", "@all", "-t user 直接接收消息的用户昵称")
	agentid := flag.Int("i", 0, "-i 0 指定agentid")
	content := flag.String("c", "Hello world", "-c 'Hello world' 指定要发送的内容")
	corpid := flag.String("p", "", "-p corpid 必须指定")
	corpsecret := flag.String("s", "", "-s corpsecret 必须指定")
	flag.Parse()

	if *corpid == "" || *corpsecret == "" {
		flag.Usage()
		return
	}

	var m send_msg = send_msg{Touser: *touser, Msgtype: "text", Agentid: *agentid, Text: map[string]string{"content": *content}}

	if *mfile != "" {
		buf, err := Parse(*mfile)
		if err != nil {
			println(err.Error())
			return
		}
		err = json.Unmarshal(buf, &m)
		if err != nil {
			println(err)
			return
		}
	}
	///-p "wx2468f5838693e123" -s "JbjkM1jYq8g3GaHjOTgj27y4n4_7Dsv4FV94I5BMRSrBsm_aTsMUVJMhGu_DFGDSF"
	token, err := Get_token(*corpid, *corpsecret)
	if err != nil {
		println(err.Error())
		return
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = Send_msg(token.Access_token, buf)
	if err != nil {
		println(err.Error())
	}
}

//发送消息.msgbody 必须是 API支持的类型
func Send_msg(Access_token string, msgbody []byte) error {
	body := bytes.NewBuffer(msgbody)
	resp, err := http.Post(sendurl+Access_token, "application/json", body)
	if resp.StatusCode != 200 {
		return requestError
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var e send_msg_error
	err = json.Unmarshal(buf, &e)
	if err != nil {
		return err
	}
	if e.Errcode != 0 && e.Errmsg != "ok" {
		return errors.New(string(buf))
	}
	return nil
}

//通过corpid 和 corpsecret 获取token
func Get_token(corpid, corpsecret string) (at access_token, err error) {
	resp, err := http.Get(get_token + corpid + "&corpsecret=" + corpsecret)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = requestError
		return
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &at)
	if at.Access_token == "" {
		err = errors.New("corpid or corpsecret error.")
	}
	return
}

func Parse(jsonpath string) ([]byte, error) {
	var zs = []byte("//")
	File, err := os.Open(jsonpath)
	if err != nil {
		return nil, err
	}
	defer File.Close()
	var buf []byte
	b := bufio.NewReader(File)
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		line = bytes.TrimSpace(line)
		if len(line) <= 0 {
			continue
		}
		index := bytes.Index(line, zs)
		if index == 0 {
			continue
		}
		if index > 0 {
			line = line[:index]
		}
		buf = append(buf, line...)
	}
	return buf, nil
}
