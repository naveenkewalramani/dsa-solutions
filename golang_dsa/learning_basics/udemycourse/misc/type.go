package main

import "fmt"

func main() {
	var a = 42
	var b = "Naveen"
	fmt.Println(a)
	fmt.Printf("%T\n", a) // Print the type of the variable
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%#b\n", a)
	fmt.Printf("%#x\n", a)
	fmt.Printf("%#o\n", a)
	fmt.Printf("%#d\n", a)
}
