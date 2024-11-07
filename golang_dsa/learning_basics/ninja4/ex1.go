package main

import "fmt"

func main() {

	// array
	var x [5]int
	fmt.Println(x)
	for i := range x {
		x[i] = i + i
	}
	fmt.Println(x)
	for i, v := range x {
		fmt.Printf("%d\t%T\t%d\t%T\n", i, i, v, v)
	}

	// slice
	a := [10]int{}
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a[i] = i * 10
	}
	fmt.Println(a)
	for i, v := range a {
		fmt.Printf("%d\t%T\t%d\t%T\n", i, i, v, v)
	}

}
