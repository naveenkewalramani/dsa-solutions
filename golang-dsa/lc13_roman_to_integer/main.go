// https://leetcode.com/problems/roman-to-integer/description/
package main

import "fmt"

func main() {
	fmt.Println(romanToInteger("X"))
	fmt.Println(romanToInteger("III"))
	fmt.Println(romanToInteger("LVIII"))
	fmt.Println(romanToInteger("MCMXCIV"))
	fmt.Println(romanToInteger("DCXXI"))

}

func romanToInteger(input string) int {
	number := 0
	index := 0
	inputLen := len(input)
	for index <= inputLen-1 {
		flag := false
		char := fmt.Sprintf("%c", input[index])
		if char == "I" {
			if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "V" {
				flag = true
				number += 4
			} else if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "X" {
				flag = true
				number += 9
			} else {
				number += 1
			}
		} else if char == "V" {
			number += 5
		} else if char == "X" {
			if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "L" {
				flag = true
				number += 40
			} else if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "C" {
				flag = true
				number += 90
			} else {
				number += 10
			}

		} else if char == "L" {
			number += 50
		} else if char == "C" {
			if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "D" {
				flag = true
				number += 400
			} else if index+1 <= inputLen-1 && fmt.Sprintf("%c", input[index+1]) == "M" {
				flag = true
				number += 900
			} else {
				number += 100
			}

		} else if char == "D" {
			number += 500
		} else {
			number += 1000
		}
		if flag {
			index += 2
		} else {
			index += 1
		}
	}
	return number
}
