package main

import "fmt"

func main() {
	for x := 0; x < 100; x += 2 {
		if x%3 == 0 || x%4 == 0 {
			fmt.Println(x)
		}
	}
}
