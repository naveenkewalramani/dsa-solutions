// recursion

package main

import "fmt"

func main() {
	fmt.Println(fact(10))
	fmt.Println(fact(19))
	fmt.Println(fact(20))
}

func fact(a int64) int64 {
	if a == 1 {
		return a
	}
	b := fact(a - 1)
	return a * b
}
