package main
import (
	"fmt"
	"sync"
)
func main(){
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 10000; j++ {
			i += 1
		}
	}()
	wg.Add(1)
	go func() {
		wg.Done()
		for j := 0; j < 10000; j++{
			i += 1
		}
	}()
	wg.Wait()
	fmt.Println(i)
}
