package main

import "fmt"

func main() {
	x := 201
	fmt.Printf("%d\t%b\t%o\t%x\n", x, x, x, x)
	fmt.Printf("%d\t%#b\t%#o\t%#x\n", x, x, x, x)
}
