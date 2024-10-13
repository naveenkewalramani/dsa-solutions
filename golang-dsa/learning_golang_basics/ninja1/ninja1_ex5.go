package main

import "fmt"

type customInt3 customInt2

func main() {
	var y customInt3
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
