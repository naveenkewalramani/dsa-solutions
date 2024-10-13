package main
import "fmt"
func main(){
	c := make(chan int, 2)
	go func(){
		c <- 1
		c <- 10
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <- c
	fmt.Println(v, ok)
}