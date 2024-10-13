package main

import "fmt"

type customInt2 int

func main() {
	var x customInt2
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
