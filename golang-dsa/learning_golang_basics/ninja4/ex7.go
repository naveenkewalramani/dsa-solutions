package main

import "fmt"

func main() {
	x := [][]string{
		{"a", "b", "C"}, {"D", "E", "f"},
	}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(v)
		for j, u := range x[i] {
			fmt.Println(j, u)
		}
		fmt.Println("Row completed")
	}
}
