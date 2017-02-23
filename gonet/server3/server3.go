package main

import (
	"bufio"
	"fmt"
	"net"
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
		reciveStr := data
		messages <- reciveStr
	}
}
func echoHandler(conns *map[string]net.Conn, messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
		fmt.Println(conns)
		for key, value := range *conns {
			fmt.Println("connection is connected from...", key)
			_, err := fmt.Fprintf(value, msg)
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

func StartServer() {
	server := ":" + "9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err, "ResolveTcpAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	conns := make(map[string]net.Conn)
	messages := make(chan string, 20)
	go echoHandler(&conns, messages)
	for {
		fmt.Println("Listening...")
		conn, err := l.Accept()
		checkError(err, "accept")
		fmt.Println("accepting..")
		conns[conn.RemoteAddr().String()] = conn
		go Handler(conn, messages)
	}
}
