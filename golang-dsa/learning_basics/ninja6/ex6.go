package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("hello i am defer")
	}()
	func() {
		fmt.Println("I am main")
	}()
}
