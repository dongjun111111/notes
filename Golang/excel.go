package main

import (
	"bufio"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"net/smtp"
	"os"
	"regexp"
	"strings"
)

func main() {
	excelFileName := "./list.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	sendList := make(map[string]string)

	for _, sheet := range xlFile.Sheets {
		curMail := ""
		for _, row := range sheet.Rows {
			cells := getCellValues(row)
			//如果行包含电子邮件，创建一个新字典项
			if isEmail, emailStr := isEmailRow(cells); isEmail {
				curMail = emailStr
			} else {
				count := 0
				for _, c := range cells {
					if len(c) > 0 {
						count++
					}
				}
				if count > 1 {
					sendList[curMail] += fmt.Sprintf("<tr><td>%s</td></tr>", strings.Join(cells, "</td><td>"))
				} else {
					sendList[curMail] += fmt.Sprintf("<tr><td colspan='%d'>%s</td></tr>", len(cells), strings.Join(cells, ""))
				}
			}
		}
	}

	sendMail(sendList)
	fmt.Print("按下回车结束")
	bufio.NewReader(os.Stdin).ReadLine()

}

func getCellValues(r *xlsx.Row) (cells []string) {
	for _, cell := range r.Cells {
		txt := strings.Replace(strings.Replace(cell.Value, "\n", "", -1), " ", "", -1)
		cells = append(cells, txt)
	}
	return
}

func isEmailRow(r []string) (isEmail bool, email string) {
	reg := regexp.MustCompile(`^[a-zA-Z_0-9.-]{1,64}@([a-zA-Z0-9-]{1,200}.){1,5}[a-zA-Z]{1,6}$`)
	for _, v := range r {
		if reg.MatchString(v) {
			return true, v
		}
	}
	return false, ""
}

func sendMail(sendList map[string]string) {

	fmt.Printf("共需要发送%d封邮件\n", len(sendList))
	index := 1
	for mail, content := range sendList {
		fmt.Printf("发送第%d封", index)
		if err := sendToMail("xxx@mybigcompany.com",
			"thesismypassword",
			"smtp.mybigcompany.com:25",
			mail,
			"工资条",
			fmt.Sprintf("<table border='2'>%s</table>", content),
			"html"); err != nil {
			fmt.Printf(" ... 发送错误(X) %s %s \n", mail, err.Error())
		} else {
			fmt.Printf(" ... 发送成功(V) %s \n", mail)
		}
		index++
	}

}

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	auth := smtp.PlainAuth("", user, password, strings.Split(host, ":")[0])
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + "Content-Type: text/" + mailtype + "; charset=UTF-8" + "\r\n\r\n" + body)
	sendto := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendto, msg)
	return err
}
