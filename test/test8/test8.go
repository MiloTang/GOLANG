package main

import (
	"fmt"
)

func init() {

}

type Show interface {
	show()
}
type FuncT func(int, int)

func main() {
	Type(`show`)
}

func show() {
	fmt.Println("just show me")
}
func Type(t interface{}) {
	fmt.Println(t)
}
