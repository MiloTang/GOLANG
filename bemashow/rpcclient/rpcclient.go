package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func init() {

}
func main() {
	//如同步的方式
	var reply int
	if client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234"); err != nil {
		fmt.Println("dialing:", err)
	} else {
		args := &Args{17, 8}
		if err = client.Call("Arith.Multiply", args, &reply); err != nil {
			fmt.Println("arith error:", err)
		} else {
			fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
		}
		//或者异步的方式
		quotient := new(Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		replyCall := <-divCall.Done
		fmt.Println(replyCall.Reply)

		//
		if address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1235"); err != nil {
			fmt.Println(err)
		} else {
			conn, _ := net.DialTCP("tcp", nil, address)
			defer conn.Close()
			client := rpc.NewClient(conn)
			defer client.Close()
			args := &Args{10, 8}
			if err = client.Call("Arith.Multiply", args, &reply); err != nil {
				fmt.Println("arith error:", err)
			} else {
				fmt.Println(reply)
			}
		}
	}
}
