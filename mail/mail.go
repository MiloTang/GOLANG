package mail

import (
	"encoding/base64"
	"fmt"
	"net"
	"time"
)

const (
	tb string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

type Sender struct {
	UserName string
	Password string
	From     string
	To       string
	Subject  string
	Text     string
	Host     string
}

func init() {

}
func NewSender(host, userName, password, to, subject, text string) *Sender {
	sender := &Sender{}
	sender.UserName = userName
	sender.Password = password
	sender.From = userName
	sender.To = to
	sender.Subject = subject
	sender.Text = text
	sender.Host = host
	return sender
}
func (this *Sender) SendMail() error {
	var (
		deadline time.Duration = 5 * time.Second
		encoding *base64.Encoding
		buf      []byte = make([]byte, 512)
		r        int
		e        error
	)
	encoding = base64.NewEncoding(tb)
	conn, err := net.Dial("tcp", this.Host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		fmt.Println(e)
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte("EHLO Juxuny\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte("AUTH LOGIN\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte(encoding.EncodeToString([]byte(this.UserName)) + "\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte(encoding.EncodeToString([]byte(this.Password)) + "\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte("MAIL FROM: <" + this.From + ">" + "\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte("RCPT TO <" + this.To + ">\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		return e
	}
	fmt.Println(string(buf[:r-1]))

	conn.Write([]byte("DATA\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte("Message-ID: <" + this.From + ">\r\n"))
	conn.Write([]byte("X-Mailer: <MMail 1.0>"))
	conn.Write([]byte("MIME-Version: 1.0"))
	conn.Write([]byte("Content-Type: text/plain"))
	conn.Write([]byte("From: <" + this.From + ">\r\n"))
	conn.Write([]byte("To: <" + this.To + ">\r\n"))
	conn.Write([]byte("Subject: " + this.Subject + "\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(this.Text))
	conn.Write([]byte("\r\n.\r\n"))

	time.Sleep(5e9)
	conn.Write([]byte("QUIT\r\n"))
	conn.SetDeadline(time.Now().Add(deadline))
	r, e = conn.Read(buf)
	if e != nil {
		fmt.Println(e)
		return e
	}
	fmt.Println(string(buf[:r-1]))

	return e
}
