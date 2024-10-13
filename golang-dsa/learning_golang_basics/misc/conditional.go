package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Divisible by 2")
		} else if i%3 == 0 {
			fmt.Println("Divisible by 3")
		} else {
			fmt.Println("Other")
		}
	}

	if x := 3; x%3 == 0 {
		fmt.Println("This is divisible by 3")
	}
}
