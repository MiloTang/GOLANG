package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func init() {

}
func handleConn(conn net.Conn) {
	var err error
	var data string
	data, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("get client data error: ", err)
		return
	}
	fmt.Printf("%v", data)
	fmt.Fprintf(conn, "hello:"+conn.RemoteAddr().String())
}
func main() {
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConn(conn)
	}
}
