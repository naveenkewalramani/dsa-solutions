package main

import (
	"fmt"
)

func main() {
	s := []int{11, 92, 83, 34, 15, 26}
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}

		}
	}
	fmt.Println(s)
}
