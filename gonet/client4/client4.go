package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var (
	List, inTE, outTE *walk.TextEdit = nil, nil, nil
	conn              *net.TCPConn   = nil
	err               error
)

func main() {
	StartClient()
	Windows()
	fmt.Fprintf(conn, "list:::"+"\n")
}
func Windows() {
	MainWindow{
		Title:   "群聊",
		MinSize: Size{600, 400},
		Layout:  HBox{},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &outTE, ReadOnly: true},
					Composite{
						MaxSize: Size{0, 50},
						Layout:  HBox{},
						Children: []Widget{
							TextEdit{AssignTo: &inTE},
							PushButton{
								Text:      "发送",
								OnClicked: ChatSend,
							},
						},
					},
				},
			},
			TextEdit{AssignTo: &List, ReadOnly: true},
		},
	}.Run()
}
func StartClient() {
	tcpAddr, e := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
	checkError(e, "ResolveTCPAddr")
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	go ReadServer()

}
func ChatSend() {
	var input string = ""
	input = inTE.Text()
	_, err := fmt.Fprintf(conn, "talk:::游客-->Say->"+input+"\n")
	if err != nil {
		fmt.Println(err.Error(), "Write")
		conn.Close()
	}
	inTE.SetText("")
}
func ReadServer() {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if checkError(err, "Connection") == false {
			conn.Close()
			fmt.Println("server is dead .....byebye")
			os.Exit(0)
		}
		s := strings.Split(data, ":::")
		if s[0] == "talk" {
			outTE.AppendText(s[1] + "\n")
		} else {
			List.SetText(s[1] + "\n")
		}

	}
}
func checkError(err error, info string) (res bool) {
	if err != nil {
		fmt.Println(info + "" + err.Error())
		return false
	} else {
		return true
	}
}
