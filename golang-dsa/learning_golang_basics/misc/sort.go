package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{11, 92, 83, 34, 15, 26}
	sort.Ints(s)
	fmt.Println(s)
}
