package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Arith int

func init() {

}
func (t *Arith) Multiply(args *Args, reply *([]string)) error {
	*reply = append(*reply, "test")
	return nil
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dial fail:", err)
	}
	args := &Args{7, 8}
	reply := make([]string, 10)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Println(reply)
}
