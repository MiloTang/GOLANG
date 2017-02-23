package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var username string = "milo"

func main() {
	fmt.Println("please input your nick name:")
	var in string
	fmt.Scanln(&in)
	username = in
	fmt.Println("you can talk now:")
	StartClient()
}
func StartClient() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "9.112.45.77:9999")
	checkError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	go chatSend(conn)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if checkError(err, "Connection") == false {
			conn.Close()
			fmt.Println("server is dead .....byebye")
			os.Exit(0)
		}
		fmt.Println(data)
	}

}
func chatSend(conn net.Conn) {
	var input string
	for {
		in := bufio.NewReader(os.Stdin)
		input, _ = in.ReadString('\n')
		if strings.Trim(input, "\r\n") == "Q" {
			fmt.Println("byebye")
			conn.Close()
			os.Exit(0)
		}
		_, err := fmt.Fprintf(conn, username+"->Say:::"+input+"\n")
		if err != nil {
			fmt.Println(err.Error(), "Write")
			conn.Close()
			break
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
