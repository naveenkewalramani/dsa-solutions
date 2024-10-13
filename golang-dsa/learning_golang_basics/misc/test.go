package main

import (
	"fmt"
	"math"
)

// argument with return type
// return type sepcify
func add(num1 int, num2 int) int {
	var result = num1 + num2
	return result
}
func sub(num1, num2 int) int {
	return num1 - num2
}

// function with multiple return type
func addSub(num1, num2 int) (int, int) {
	result1 := add(num1, num2)
	result2 := sub(num2, num1)
	return result1, result2
}

func main() {
	var name = "Hello"
	// creation and assignment
	var num1 = 12
	// creation and assignment with data type
	var num3 int = 23
	var num2 float64 = 9
	// shorthand notation of initalization and declration and calling method returning multiple var
	resultAdd, resultSub := addSub(num1, num3)
	// creation of constants
	const defaultNum = 242
	// for loop(only one loop in goLang)
	for i := 1; i < 5; i++ {
		// fmt.Println("12")
	}

	fmt.Println(name)
	fmt.Println(num3)
	fmt.Println(resultAdd)
	fmt.Println(resultSub)
	fmt.Println(defaultNum)
	// var number int
	// fmt.Scanln(&number)
	// fmt.Println(number)
	fmt.Println(math.Sqrt(num2))
}
