package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var (
	List, inTE, outTE *walk.TextEdit = nil, nil, nil
	conn              *net.TCPConn   = nil
	err               error
	lists             string = ""
)

func main() {
	StartClient()
	Windows()
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
	tcpAddr, e := net.ResolveTCPAddr("tcp4", "9.112.45.77:9999")
	checkError(e, "ResolveTCPAddr")
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	go ReadServer()

}
func ChatSend() {
	var input string = ""
	input = inTE.Text()
	_, err := fmt.Fprintf(conn, input+"\n")
	if err != nil {
		fmt.Println(err.Error(), "Write")
		conn.Close()
	}
	inTE.SetText("")
}
func Update() {

}
func ReadServer() {
	for {
		data, e := bufio.NewReader(conn).ReadString('\n')
		if checkError(e, "Connection") == false {
			conn.Close()
			fmt.Println("server is dead .....byebye")
			os.Exit(0)
		}
<<<<<<< HEAD
		outTE.AppendText(data + "\n")
=======
		if data == "diu" {
			s := strings.Split(data, ":::")
			if s[0] == "talk" {
				outTE.AppendText(s[1] + "\n")
			} else {
				List.SetText("")
				for i, value := range s {
					if i == 0 {
						continue
					} else {
						List.AppendText(value + "\n")
					}

				}

			}
		}

>>>>>>> d16a6de7dee06f8f235393a4d4769e868047a746
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
