package main

import "fmt"

func main(){
	var a, b int
	var c string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	switch c {
	case "+":
		fmt.Println(a+b)
	case "-":
		fmt.Println(a-b)
	case "*":
		fmt.Println(a*b)
	case "/":
		fmt.Println(a/b)
	}
}
