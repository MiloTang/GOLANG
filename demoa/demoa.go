package main

import (
	"demob"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(5400))
	}
	fmt.Println("The time is", time.Now())
	rand.Seed(1)
	fmt.Println("random number", rand.Int())
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Pi)
	tot := demob.Add(11, 11)
	fmt.Println(tot)
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
