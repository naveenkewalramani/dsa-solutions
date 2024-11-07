package main

import "fmt"

func main() {
	unbufferedChannel()
	bufferedChannel()
	continueReadingUnBufferedChannel()
	continueReadingUnBufferedChannel2()
}

func unbufferedChannel() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 10
	}(ch)
	val := <-ch
	fmt.Println(val)
}
func bufferedChannel() {
	ch := make(chan int, 1)
	func(ch chan int) {
		ch <- 10
	}(ch)
	val := <-ch
	fmt.Println(val)
}

func continueReadingUnBufferedChannel() {
	ch := make(chan int)
	quit := make(chan bool)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(quit)
	}(ch)
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-quit:
			fmt.Println("Finish reading from channel 1")
			return
		}
	}
}

func continueReadingUnBufferedChannel2() {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Finish reading from channel 2")
				return
			}
			fmt.Println(val * 2)
		}
	}
}
