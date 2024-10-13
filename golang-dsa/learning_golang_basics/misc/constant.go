package main

import "fmt"

// A is a constant
const A = 20

// B, C, D are constants decalred together but no type defined(untyped constants)
const (
	B = 21
	C = 21.902
	D = "Twenty One"
)

// E,F are constants decalred together with type defined(typed constants)
const (
	E int     = 21
	F float64 = 21.9090
)

func main() {
	fmt.Println(A)
	fmt.Printf("%T\n", A)
	fmt.Println(B)
	fmt.Printf("%T\n", B)
	fmt.Println(C)
	fmt.Printf("%T\n", C)
	fmt.Println(D)
	fmt.Printf("%T\n", D)
	fmt.Println(E)
	fmt.Printf("%T\n", E)
	fmt.Println(F)
	fmt.Printf("%T\n", F)

}
