package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

func SendEmailWithAttachment(user, passwd, host, to, subject string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, passwd, hp[0])
	buffer := bytes.NewBuffer(nil)

	boudary := "THIS_IS_BOUNDARY_JUST_MAKE_YOURS"
	header := fmt.Sprintf("To:%s\r\n"+
		"From:%s\r\n"+
		"Subject:%s\r\n"+
		"Content-Type:multipart/mixed;Boundary=\"%s\"\r\n"+
		"Mime-Version:1.0\r\n"+
		"Date:%s\r\n", to, user, subject, boudary, time.Now().String())
	buffer.WriteString(header)
	fmt.Print(header)

	msg1 := "\r\n\r\n--" + boudary + "\r\n" + "Content-Type:text/plain;charset=utf-8\r\n\r\n这是正文啊\r\n"

	buffer.WriteString(msg1)
	fmt.Print(msg1)

	msg2 := fmt.Sprintf(
		"\r\n--%s\r\n"+
			"Content-Transfer-Encoding: base64\r\n"+
			"Content-Disposition: attachment;\r\n"+
			"Content-Type:image/jpg;name=\"test.jpg\"\r\n", boudary)
	buffer.WriteString(msg2)
	fmt.Print(msg2)

	attachmentBytes, err := ioutil.ReadFile("./test.jpg")
	if err != nil {
		fmt.Println("ReadFile ./test.jpg Error : " + err.Error())
		return err
	}
	b := make([]byte, base64.StdEncoding.EncodedLen(len(attachmentBytes)))
	base64.StdEncoding.Encode(b, attachmentBytes)
	buffer.WriteString("\r\n")
	fmt.Print("\r\n")
	fmt.Print("图片base64编码")
	for i, l := 0, len(b); i < l; i++ {
		buffer.WriteByte(b[i])
		if (i+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}

	buffer.WriteString("\r\n--" + boudary + "--")
	fmt.Print("\r\n--" + boudary + "--")

	sendto := strings.Split(to, ";")
	err = smtp.SendMail(host, auth, user, sendto, buffer.Bytes())

	return err
}
