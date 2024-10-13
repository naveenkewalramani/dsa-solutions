package main

import "fmt"

func main() {
	x, y := testMethod(10, 20)
	fmt.Println(x)
	fmt.Println(y)
	variadicMethod(2, 2, 3, 3, 4, 2, 5, 5, 5, 5, 5)
}

func testMethod(a, b int) (int, int) {
	return (a + b), (a - b)
}

// similar to *args
func variadicMethod(x ...int) {
	fmt.Println(x)
}
