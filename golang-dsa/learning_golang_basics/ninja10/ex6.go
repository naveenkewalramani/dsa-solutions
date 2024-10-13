package main

import "fmt"

func main(){
	c := make(chan int)
	putToChannel(c)
	getFromChannel(c)
}

func putToChannel(c chan<- int){
	go func(){
		for j:= 1;j<= 100; j++{
			c <- j
		}
		close(c)
	}()
}

func getFromChannel(c <-chan int){
	for j :=1; j<=100;j++{
		fmt.Println(<-c)
	}
}