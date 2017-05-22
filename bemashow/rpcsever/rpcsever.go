package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
	if l, e := net.Listen("tcp", ":8888"); e != nil {
		log.Fatal("listen error:", e)
	} else {
		go http.Serve(l, nil)
	}

	newsever := rpc.NewServer()
	newsever.Register(new(Arith))
	if l, e := net.Listen("tcp", ":8889"); e != nil {
		log.Fatal("listen error:", e)
	} else {
		newsever.Accept(l)
		newsever.HandleHTTP("/a", "/b")
	}
}
