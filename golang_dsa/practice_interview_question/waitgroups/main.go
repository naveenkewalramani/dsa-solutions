package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printNum(i)
	}
	wg.Wait()
}

func printNum(i int) {
	fmt.Println(i)
	wg.Done()
}
