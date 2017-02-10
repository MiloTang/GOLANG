package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch  chan int
	str string
)

func init() {
	ch = make(chan int)
	str = "hello milo"
}
func main() {
	go say("world")
	say("hello")
	fmt.Println("---------------1")
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	fmt.Println("---------------2")
	c2 := make(chan int, 3)
	c2 <- 1
	c2 <- 2
	c2 <- 3
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println("---------------3")
	c3 := make(chan int, 2)
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}
	fmt.Println("---------------4")
	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(<-c4)
		}
		quit <- 0
	}()
	fibonacci2(c4, quit)
	fmt.Println("---------------5")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
	}
	for i := 0; i < 10; i++ {
		go done(&wg, i)
	}
	wg.Wait()
	fmt.Println("---------------6")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick. ")
		case <-boom:
			fmt.Println("BOOM!")
			fmt.Println("---------------end")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
func add(wg sync.WaitGroup) {
	wg.Add(1)
}
func done(wg *sync.WaitGroup, i int) {
	fmt.Println("sync", i)
	wg.Done()
}
func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
