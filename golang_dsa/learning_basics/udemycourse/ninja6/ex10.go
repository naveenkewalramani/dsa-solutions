//closure
package main

import "fmt"

func main() {
	x := 0
	y := 1
	fmt.Println("Value of x outside the anon. function: ", x)
	fmt.Println("Value of y outside the anon. function: ", y)
	func() {
		x := 20
		y := 30
		fmt.Println("Value of x inside the anon. function: ", x)
		fmt.Println("Value of y inside the anon. function: ", y)
	}()
	fmt.Println("Value of x outside the anon. function: ", x)
	fmt.Println("Value of y outside the anon. function: ", y)
}
