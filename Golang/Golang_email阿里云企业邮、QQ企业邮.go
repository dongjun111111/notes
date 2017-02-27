package email

import (
	"net/smtp"
	"strings"
)

//qq企业邮箱
func EmailQQ(title, content, toWho string) error {
	host := "smtp.exmail.qq.com:25"
	to := strings.Split(toWho, ";")
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + toWho + "\r\nFrom: Jason-Test \r\nSubject:" + title + "\r\n" + content_type + "\r\n\r\n" + content)
	err := smtp.SendMail(host, smtp.PlainAuth("", "from@from.com", "from", "smtp.exmail.qq.com"), "from@from.com", to, []byte(msg))
	return err
}

//阿里云企业邮箱
func EmailAli(title string, content *string, toWho string) error {
	toWho += ";to@to.com"
	host := "smtp.mxhichina.com:25"
	to := strings.Split(toWho, ";")
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + toWho + "\r\nFrom: Jason-Test \r\nSubject:" + title + "\r\n" + content_type + "\r\n\r\n" + *content)
	err := smtp.SendMail(host, smtp.PlainAuth("", "from@from.com", "from", "smtp.mxhichina.com"), "from@from.com", to, []byte(msg))
	return err
}
