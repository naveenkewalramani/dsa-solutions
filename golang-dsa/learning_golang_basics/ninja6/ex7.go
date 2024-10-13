package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Function assigned to var")
	}
	a()
	b := test()
	b()
}

func test() func() {
	return func() {
		fmt.Println("Function returned from test() method")
	}
}
