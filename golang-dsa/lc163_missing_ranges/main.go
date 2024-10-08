package main

import "fmt"

func main() {
	fmt.Println(missingRange([]int{}, 10, 20))
	fmt.Println(missingRange([]int{1, 2, 3, 4, 6}, -4, 20))
}

func missingRange(nums []int, lower, upper int) [][]int {
	outputArr := make([][]int, 0)
	n := len(nums)
	if n == 0 {
		outputArr = append(outputArr, []int{lower, upper})
		return outputArr
	}
	if lower < 0 {
		outputArr = append(outputArr, []int{lower, nums[0] - 1})
	}
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			outputArr = append(outputArr, []int{nums[i] + 1, nums[i+1] - 1})
		}
	}
	if nums[n-1] < upper {
		outputArr = append(outputArr, []int{nums[n-1] + 1, upper})
	}

	return outputArr
}
