package main

import "fmt"

type person struct {
	name string
	age  int
	addr string
}

func main() {
	p1 := person{
		name: "Naveen",
		age:  24,
		addr: "E-390 Ajmer",
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age, p1.addr)
	fmt.Println(&p1.name, &p1.age, &p1.addr)
	changeMe(&p1)
	fmt.Println(p1.name, p1.age, p1.addr)
}

func changeMe(p *person) {
	p.addr = "E-391 Ajmer"
}
