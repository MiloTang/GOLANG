package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var str string

func init() {
	str = "init str"
}
func main() {
	flag.Parse()
	fmt.Println(os.Args, flag.Args())
	return
	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	host := os.Args[1]
	port := flag.Arg(1)
	fmt.Println(port)
	addr := net.ParseIP(host)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
