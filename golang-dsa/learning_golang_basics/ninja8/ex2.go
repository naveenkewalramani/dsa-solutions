package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{21, 32, 153, 54, 25, 67}
	y := []string{"bbbw", "eueuq", "qwsq", "wwqwdq", "aaas", "pope", "ew"}
	fmt.Println(x)
	fmt.Println(y)
	sort.Ints(x)
	sort.Strings(y)
	fmt.Println(x)
	fmt.Println(y)
}
