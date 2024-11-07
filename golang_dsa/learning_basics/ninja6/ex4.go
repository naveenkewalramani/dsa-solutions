package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Naveen",
		last:  "Kewalramani",
		age:   20,
	}
	fmt.Println(p)
	p.speak()
}

func (p person) speak() {
	fmt.Println("Name is : ", p.first, p.last, " and age : ", p.age)
}
