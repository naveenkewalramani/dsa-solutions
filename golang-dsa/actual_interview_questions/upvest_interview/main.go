package main

// build a publisher, and multiple subscriber consuming messaging from same
import (
	"fmt"
	"sync"
)

var (
	subscribers = []chan int{make(chan int, 5), make(chan int, 5), make(chan int, 5)}
)

func broker(element int) {
	for _, ch := range subscribers {
		ch <- element
	}
}

func publisher(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Publishing msg with id:", i)
		broker(i)
	}
	for _, x := range subscribers {
		close(x)
	}
}

func subscriber(wg *sync.WaitGroup, subID int) {
	defer wg.Done()
	ch := subscribers[subID]
	for msg := range ch {
		fmt.Printf("Receive msg from channel, msg: %+v, id: %d \n", msg, subID)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go publisher(&wg)

	numOfSubs := 3
	for i := 0; i < numOfSubs; i++ {
		wg.Add(1)
		go subscriber(&wg, i)
	}
	wg.Wait()
}
