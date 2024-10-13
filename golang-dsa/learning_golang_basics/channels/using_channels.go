package main

import "fmt"

func main() {
	c := make(chan int) // channels are refrence type
	go sender(c)        //sender
	receiver(c)         //receiver
	fmt.Println("Function exit from here")
}

func sender(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}

func receiver(channel chan int) {
	for v := range channel {
		fmt.Println(v)
	}
}
