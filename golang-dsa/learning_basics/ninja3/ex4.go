package main

import "fmt"

func main() {
	x := 1997
	y := 2020

	for {
		fmt.Println(x)
		x++
		if x > y {
			break
		}
	}
}
