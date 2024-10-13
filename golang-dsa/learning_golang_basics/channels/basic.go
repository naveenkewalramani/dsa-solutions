package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello Brother1"
	}()
	go func() {
		channel <- "Hello Brother2"
	}()
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	channel1 := make(chan int, 2)
	channel1 <- 32
	fmt.Println(<-channel1)
}
