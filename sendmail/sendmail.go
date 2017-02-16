package main

import (
	"fmt"
	"mail"
)

func main() {
	sender := mail.NewSender("pop.qq.com:587",
		"tangguang1986@163.com",
		"*******",
		"mainframe_fb@163.com",
		"hello mail",
		"Test mail")
	e := sender.SendMail()
	if e != nil {
		fmt.Println(e)
	}

}
