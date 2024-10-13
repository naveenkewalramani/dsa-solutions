package main

import "fmt"

var a int

type customInt int

var b customInt

func main() {
	a = 20
	b = 21
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
