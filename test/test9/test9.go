package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

const N int = int(unsafe.Sizeof(0))

func init() {

}

type Speaker interface {
	Say()
}
type Human struct {
	name string
}
type Dog struct {
	name string
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func main() {
	fmt.Println(GetFunctionName(Show))
	h := new(Human)
	h.name = "Ben"
	d := new(Dog)
	d.name = "Tom"
	Say(h)
	Say(d)
	ch := make(chan int)
	go pump(ch) //异步执行pump
	go suck(ch) //异步执行suck
	x := 0x1234
	p := unsafe.Pointer(&x)
	fmt.Println((*[N]byte)(p))
	p2 := (*[N]byte)(p)
	if p2[0] == 0 {
		fmt.Println("本机器：大端")
	} else {
		fmt.Println("本机器：小端")
	}

}

func Say(s Speaker) {
	s.Say()
}

func (h *Human) Say() {
	fmt.Println(h.name + " Say Hello Go")
}

func (d *Dog) Say() {
	fmt.Println(d.name + " Say Hello Go")
}
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
func Show() {
	fmt.Println("show")
}
