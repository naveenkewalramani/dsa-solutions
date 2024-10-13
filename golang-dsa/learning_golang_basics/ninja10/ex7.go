package main

import "fmt"

func main(){
	c := make(chan int)
	for i:=1; i<=100;i+=10{
		addToChannel(c, i)
	}
	fetchFromChannel(c)
}

func addToChannel(c chan<- int, i int){
	go func(){
		for j:= 1; j<=10;j++ {
			c <- i + j
		}
	}()
}

func fetchFromChannel(c <-chan int){
	for i:=1;i<=100;i++{
		fmt.Println(<-c)
	}
}