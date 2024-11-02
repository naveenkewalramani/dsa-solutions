// passing a func to another func as an argument is called as callback

package main

import "fmt"

func main() {
	a := func() int {
		return 34
	}
	c(a)
}
func c(a func() int) {
	fmt.Println(a())
}
