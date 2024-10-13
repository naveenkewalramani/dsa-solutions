package main

import (
	"fmt"
)

func main() {
	s := []int{11, 92, 83, 34, 15, 26}
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if i != min {
			s[min], s[i] = s[i], s[min]
		}
	}
	fmt.Println(s)
}
