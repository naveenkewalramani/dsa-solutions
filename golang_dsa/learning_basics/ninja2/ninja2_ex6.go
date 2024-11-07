package main

import (
	"fmt"
)

const (
	a = 2020 - iota
	b = 2020 - iota
	c = 2020 - iota
	d = 2020 - iota
)

const (
	_  = iota
	a1 = 2020 + iota
	b1 = 2020 + iota
	c1 = 2020 + iota
	d1 = 2020 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
}
