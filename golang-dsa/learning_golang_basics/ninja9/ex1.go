package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("In Main Go Routine")
	wg.Add(2)
	go c()
	go a()
	wg.Wait()
	fmt.Println("In Main Go Complete")
}

func c() {
	fmt.Println("This method is c")
	wg.Done()
}
func a() {
	fmt.Println("This method is a")
	wg.Done()
}
