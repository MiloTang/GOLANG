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
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("get client data error: ", err)
		return
	}
	fmt.Printf("%v", data)
	fmt.Fprintf(conn, "hello client\n")
	conn.Close()
}
func main() {
	ln, err := net.Listen("tcp", "localhost:6001")
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
