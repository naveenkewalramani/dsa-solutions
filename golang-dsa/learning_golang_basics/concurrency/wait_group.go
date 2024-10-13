package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// when main func is completed, all go routine shuts down even if another go routine is not completed
func main() {
	// Basics Specs
	fmt.Println("OS : ", runtime.GOOS)
	fmt.Println("ARCH : ", runtime.GOARCH)
	fmt.Println("CPU : ", runtime.NumCPU())
	fmt.Println("Go Call : ", runtime.NumCgoCall())
	fmt.Println("Go Routine : ", runtime.NumGoroutine())

	wg.Add(1)
	go a()
	fmt.Println("Go Routine : ", runtime.NumGoroutine())
	b()
	wg.Wait()
	fmt.Println("Go Routine : ", runtime.NumGoroutine())
}
func a() {
	for i := 0; i < 20; i++ {
		fmt.Println("i : ", i)
	}
	wg.Done()
}
func b() {
	for j := 0; j < 20; j++ {
		fmt.Println("j : ", j)
	}
}
