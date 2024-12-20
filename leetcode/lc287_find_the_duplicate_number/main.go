package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 3}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 3, 3}))
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
