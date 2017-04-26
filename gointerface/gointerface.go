package main

import (
	"fmt"
)

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

func main() {
	h := new(Human)
	h.name = "Ben"
	d := new(Dog)
	d.name = "Tom"
	Say(h)
	Say(d)
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
