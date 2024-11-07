package main

import "fmt"

func main(){
	// Channel block and therefore separate go routine to write and reading from main
	c := make(chan int)
	go func(c chan int) {
		c <- 73
	}(c)
	fmt.Println(<-c)

	// Ways to read using unbuffered channel
	b := make(chan int, 1)
	b <- 48
	fmt.Println(<-b)

}