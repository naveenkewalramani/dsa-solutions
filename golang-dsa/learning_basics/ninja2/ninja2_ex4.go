package main

import (
	"fmt"
)

func main() {
	a := 4542
	fmt.Printf("%d\t%#b\t%#x\t%#o\n", a, a, a, a)
	b := a << 1
	c := a >> 1
	fmt.Printf("%d\t%#b\t%#x\t%#o\n", b, b, b, b)
	fmt.Printf("%d\t%#b\t%#x\t%#o\n", c, c, c, c)
}
