package main

import "fmt"

func main() {
	fmt.Println("Begin")
	defer a()
	fmt.Println("End")
}

func a() {
	fmt.Println("Defer")
}
