package main

import (
	"fmt"
	"sync"
)

var (
	sharedMemoryVariable         = 10
	sharedMemoryVariableWithLock = 10
	wg                           sync.WaitGroup
	lock                         = sync.Mutex{}
)

func main() {
	causingRaceCondition()
	notCausingRaceCondition()
}

func causingRaceCondition() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go updateMemoryProcess(i)
	}
	wg.Wait()
	fmt.Println(sharedMemoryVariable)
}

func updateMemoryProcess(i int) {
	sharedMemoryVariable += i
	wg.Done()
}

func notCausingRaceCondition() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go updateMemoryWithLockProcess(i)
	}
	wg.Wait()
	fmt.Println(sharedMemoryVariableWithLock)
}
func updateMemoryWithLockProcess(i int) {
	lock.Lock()
	sharedMemoryVariableWithLock += i
	lock.Unlock()
	wg.Done()
}
