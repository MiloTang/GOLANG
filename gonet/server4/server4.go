package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var (
	num int = 0
)

func main() {
	StartServer()
}
func checkError(err error, info string) (res bool) {
	if err != nil {
		fmt.Println(info + "" + err.Error())
		return false
	} else {
		return true
	}
}

func Handler(conn net.Conn, messages chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if checkError(err, "connection") == false {
			conn.Close()
			break
		}
		s := strings.Split(conn.RemoteAddr().String(), ":")
		reciveStr := s[0] + "->" + data
		messages <- reciveStr
	}
}
func echoHandler(conns *map[string]net.Conn, messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
		for key, value := range *conns {
			fmt.Println("connection is connected from...", key)
			_, err := fmt.Fprintf(value, msg)
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
				if num != len(*conns) {
					Lists(conns)
					num = len(*conns)
				}
			}
		}
	}
}
func Lists(conns *map[string]net.Conn) {
	var lists string = ""
	for key, _ := range *conns {
		lists = lists + " " + key
	}
	for _, value := range *conns {
<<<<<<< HEAD
		fmt.Println(value, "list:::"+lists)
=======
		fmt.Fprintf(value, "list:::"+lists+"\n")
>>>>>>> d16a6de7dee06f8f235393a4d4769e868047a746
	}
	lists = ""
}
func StartServer() {
	server := ":" + "9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err, "ResolveTcpAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	conns := make(map[string]net.Conn)
	messages := make(chan string, 50)
	go echoHandler(&conns, messages)
	for {
		fmt.Println("Listening...")
		conn, err := l.Accept()
		checkError(err, "accept")
		fmt.Println("accepting..")
		conns[conn.RemoteAddr().String()] = conn
		if num != len(conns) {
			Lists(&conns)
			num = len(conns)
		}
		go Handler(conn, messages)
	}
}
