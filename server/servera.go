package main

import (
	"log"
	"net"
)

func init() {

}

const BUF_SIZE = 50

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, e := l.Accept()
		if e != nil {
			log.Println(e)
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, BUF_SIZE)
	n, err := conn.Read(buf)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}
	log.Printf("\n已接收：%d个字节，数据是：'%s'\n", n, string(buf))
}
