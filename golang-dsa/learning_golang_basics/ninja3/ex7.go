package main

import "fmt"

func main() {
	x := 7
	if x%2 == 0 {
		fmt.Println("divisible by 2")
	} else if x%7 == 0 {
		fmt.Println("divisible by 7")
	} else {
		fmt.Println("divisible by 2 or 7")
	}
}
