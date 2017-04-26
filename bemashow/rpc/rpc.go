package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func init() {

}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	if l, e := net.Listen("tcp", ":1234"); e != nil {
		log.Fatal("listen error:", e)
	} else {
		go http.Serve(l, nil)
	}
	time.Sleep(5 * time.Second)
	//比如同步的方式
	if client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234"); err != nil {
		log.Fatal("dialing:", err)
	} else {
		args := &Args{7, 8}
		var reply int
		if err = client.Call("Arith.Multiply", args, &reply); err != nil {
			log.Fatal("arith error:", err)
		} else {
			fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
		}
		//或者异步的方式
		quotient := new(Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		replyCall := <-divCall.Done
		fmt.Println(replyCall.Reply)
	}

}
