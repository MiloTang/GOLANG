package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func init() {

}
func handleConn(conn net.Conn) {
	var err error
	var data string
	var file *os.File
	data, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("get client data error: ", err)
		conn.Close()
		return
	}
	fmt.Printf("%v", data)
	s := strings.Split(conn.RemoteAddr().String(), ":")
	_, err = os.Stat(s[0])
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(s[0])
			if err != nil {
				fmt.Println("get client data error: ", err)
				conn.Close()
				return
			}
		} else {
			fmt.Println("get client data error: ", err)
			conn.Close()
			return
		}
	} else {
		file, err = os.OpenFile(s[0], os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("get client data error: ", err)
			conn.Close()
			return
		}
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("get client data error: ", err)
		conn.Close()
		return
	}
	//	fmt.Fprintf(conn, "hello client\n")
	//fmt.Println(conn.RemoteAddr())
	conn.Close()
}
func main() {
	ln, err := net.Listen("tcp", ":6001")
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
