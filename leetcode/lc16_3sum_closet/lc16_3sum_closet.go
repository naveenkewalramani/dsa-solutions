// https://leetcode.com/problems/3sum-closest/description/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Test Case 1 output: %d\n", find3SumCloset([]int{1, 2, -1, 4}, 1))
	fmt.Printf("Test Case 1 output: %d\n", find3SumCloset([]int{0, 0, 0, 0}, 0))
}

func find3SumCloset(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if absolute(target-sum) < absolute(target-result) {
				result = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return result
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
