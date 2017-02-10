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
func main() {
	for {
		fmt.Println("please enter:")
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Fprintf(conn, line)
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", data)
	}

}
