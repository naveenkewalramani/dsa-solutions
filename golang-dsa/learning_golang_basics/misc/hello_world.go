package main

import "fmt"

func main() {
	// can pass as many arguments as needed
	var byte, error = fmt.Println(22, 34, 55, 664, 343)
	fmt.Println(byte)
	fmt.Println(error)

	shortDeclare := 20
	fmt.Println(shortDeclare)
	shortDeclare = 40
	fmt.Println(shortDeclare)
}
