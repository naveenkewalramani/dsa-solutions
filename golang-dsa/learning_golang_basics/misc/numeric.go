package main

import "fmt"

var m int
var n float64

func main() {
	m := 42
	n := 43.892882
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
