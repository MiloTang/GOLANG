package main

import (
	"fmt"
	"net"
	"os"
)

func init() {

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
	buf := make([]byte, 1024)
	for {
		lenght, err := conn.Read(buf)
		if checkError(err, "connection") == false {
			conn.Close()
			break
		}
		reciveStr := string(buf)
		if lenght > 0 {
			buf[lenght] = 0
		}
		messages <- reciveStr
	}
}
func echoHandler(conns *map[string]net.Conn, messages chan string) {
	for {
		msg := <-messages
		fmt.Println(msg)
		for key, value := range *conns {
			fmt.Println("connection is connected from...", key)
			_, err := value.Write([]byte(msg))
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

func StartServer(port string) {
	server := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err, "ResolveTcpAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)
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
func chatSend(conn net.Conn) {
	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("byebye")
			conn.Close()
			os.Exit(0)
		}
		lens, err := conn.Write([]byte(username + "Say:::" + input))
		fmt.Println(lens)
		if err != nil {
			fmt.Println(err.Error(), "Write")
			conn.Close()
			break
		}
	}
}
func StartClient(tcpaddr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	checkError(err, "ResolveTCPAddr")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	go chatSend(conn)
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			fmt.Println("server is dead .....byebye")
			os.Exit(0)
		}
		fmt.Println(string(buf))
	}

}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("wrong pare")
		os.Exit(0)
	}
	if os.Args[1] == "server" && len(os.Args) == 3 {
		StartServer(os.Args[2])
	} else if os.Args[1] == "client" && len(os.Args) == 3 {
		StartClient(os.Args[2])
	} else {
		fmt.Println("error param")
		os.Exit(0)
	}

}
