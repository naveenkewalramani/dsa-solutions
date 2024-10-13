package main

import "fmt"

type person struct {
	name string
}
type secretAgent struct {
	person
	code  string
	grade int
}

func main() {
	agent1 := secretAgent{
		person: person{name: "Naveen"},
		code:   "UISSS",
		grade:  7,
	}
	fmt.Println(agent1)
	name(agent1)
	agent1.speak()
}

// function with any receiver
func name(a secretAgent) {
	fmt.Println(a.person.name)
}

// method with receiver
func (a secretAgent) speak() {
	fmt.Println(a)
}
