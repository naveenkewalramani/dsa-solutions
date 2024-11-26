package main

import (
	"fmt"
	"sort"
)

func main() {
	inputArr := []int{-1, 0, 1, 2, -1, -4}
	response := make([][]int, 0)
	n := len(inputArr)
	sort.Ints(inputArr)
	for i := 0; i < n-2; i++ {
		if i > 0 && inputArr[i] == inputArr[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			total := inputArr[i] + inputArr[left] + inputArr[right]
			if total == 0 {
				response = append(response, []int{inputArr[i], inputArr[left], inputArr[right]})
				left++
				right--
				for left < right && inputArr[left] == inputArr[left-1] {
					left++
				}
				for left < right && inputArr[right] == inputArr[right+1] {
					right--
				}
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}
	fmt.Println(response)
}
