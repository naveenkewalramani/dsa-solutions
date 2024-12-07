package main
import "fmt"
func main(){
	cs := make(chan int)
	go func(){
		cs <- 743
	}()
	fmt.Println(<-cs)
	fmt.Println("-----\n")
	fmt.Println("cs\t%T\n", cs)
}