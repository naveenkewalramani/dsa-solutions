package main

import "fmt"

type Custom int

var x Custom

func main() {
	a()   //function
	x.b() //method
	// anonymous method
	func(m, n int) {
		fmt.Println(m, n)
	}(1, 3)
}

func a() {
	fmt.Println("Normal function")
}

func (x Custom) b() {
	fmt.Println("Method - function defined on instance of an object")
}
