package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("This is true")
	case false:
		fmt.Println("This is false")
	default:
		fmt.Println("This is default case")
	}
}
