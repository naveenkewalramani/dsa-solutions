package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	fmt.Println("Value of counter intially : ", counter)
	wg.Add(100)
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			tmp := counter //read
			tmp++          //increment
			counter = tmp  //write back
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of counter is :", counter)
}
