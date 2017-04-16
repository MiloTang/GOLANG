package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

var (
	x int  = 999
	y bool = false
)

func rand_generator_2() chan int {
	out := make(chan int)
	go func() {
		for {
			out <- rand.Int()
		}
	}()
	return out
}
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("My favorite number is", r.Intn(100))
	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var c int = 110
	var d = 120
	e := 130
	println(c, d, e)
	var g, f = 123, "hello"
	i, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h, i)
	rand_service_handler := rand_generator_2()
	fmt.Printf("%d\n", <-rand_service_handler)
}
