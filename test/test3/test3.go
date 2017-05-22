package main

import "fmt"

func demoInputAndOutput(channel chan string) {
	channel <- "what is my input?"
	input := <-channel
	channel <- fmt.Sprintf("input is: %s", input)
}

func main() {
	channel := make(chan string)
	go demoInputAndOutput(channel)
	fmt.Println(<-channel)
	channel <- "42"
	fmt.Println(<-channel)
}
