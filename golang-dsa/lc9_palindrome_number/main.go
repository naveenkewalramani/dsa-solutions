package main

import "fmt"

// https://leetcode.com/problems/palindrome-number/description/

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1))
}

func isPalindrome(input int) bool {
	if input < 0 {
		return false
	}
	numbers := []int{}
	for input != 0 {
		numbers = append(numbers, input%10)
		input = input / 10
	}
	left := 0
	right := len(numbers) - 1
	for left <= right {
		if numbers[left] != numbers[right] {
			return false
		}
		left += 1
		right -= 1
	}

	return true
}
