package main

import "fmt"

func main() {
	s1 := foo([]int{1, 2, 3, 4, 5, 6, 7}...)
	s2 := bar([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(s1, s2)
}

//unfurl slice
func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
