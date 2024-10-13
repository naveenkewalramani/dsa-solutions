package main

import (
	"fmt"
)

func main(){
	a := 10
	b := 20
	result := custom_package.AddToNumbers(a, b)
	fmt.Println(result)
}

