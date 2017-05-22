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
func main() {

	for {
		conn, err := net.Dial("tcp", "127.0.0.1:6001")
		defer conn.Close()
		fmt.Println("please enter:")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		trimmedline := strings.Trim(line, "\r\n")
		if trimmedline == "Q" {
			return
		} else {
			fmt.Fprintf(conn, line)
		}
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", data)
	}

}
