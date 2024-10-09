package main

import "fmt"

func main() {
	input := []int{0, 0, 0, 1, 1, 2, 2, 2, 3, 4, 5, 5}
	x := removeDuplicateElement(input)
	fmt.Println(input[:x])
}

func removeDuplicateElement(input []int) int {
	insertAt := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] != input[i] {
			input[insertAt] = input[i]
			insertAt += 1
		}
	}
	return insertAt
}
