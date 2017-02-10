package main

import (
	"fmt"
	"runtime"
)

var i int

func init() {

}
func main() {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	go print(0)
	fmt.Println("=====", i)
}

func print(num int) {
	for {
		fmt.Println(i)
		i += 1
	}

}
