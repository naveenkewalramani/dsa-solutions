package main

import "fmt"

// UntypedString is a untyped string constant
const UntypedString = "Naveen"

// TypedString is a typed string constant
const TypedString string = "Naveen1"

func main() {
	fmt.Println(UntypedString)
	fmt.Printf("%T\n", UntypedString)
	fmt.Println(TypedString)
	fmt.Printf("%T\n", TypedString)
}
