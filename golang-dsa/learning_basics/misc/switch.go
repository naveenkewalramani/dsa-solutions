package main

import "fmt"

func main() {
	x := 1
	switch {
	case x == 3:
		fmt.Println("this is false")
	case x == 2:
		fmt.Println("this is true")
	default:
		fmt.Println("default case")
	}
}
