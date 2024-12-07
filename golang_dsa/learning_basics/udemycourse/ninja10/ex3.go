package main

import (
	"fmt"
)

func main(){
	cs := make(chan int, 12)
	for i:= 0; i< 10; i++ {
		cs <- i
	}
	for _ = range cs{
		fmt.Println(<-cs)
	}
}
