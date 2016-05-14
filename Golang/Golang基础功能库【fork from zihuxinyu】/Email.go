package Library

import (
	"fmt"
	"net/smtp"
	"strings"
)
type SmtpConfig struct {
	Username string
	Password string
	Host string
	Addr string
}
// send mail
func SendMail(subject string, message string, from string, to []string, smtpConfig SmtpConfig, isHtml bool) error {
	auth := smtp.PlainAuth(
		"",
		smtpConfig.Username,
		smtpConfig.Password,
		smtpConfig.Host,
	)
	contentType := "text/plain"
	if isHtml {
		contentType = "text/html"
	}
	msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", strings.Join(to, ";"), from, subject, contentType, message)
	return smtp.SendMail(smtpConfig.Addr, auth, from, to, []byte(msg))
}
