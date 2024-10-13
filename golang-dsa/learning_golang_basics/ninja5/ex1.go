package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	x := person{
		firstName: "Naveen1",
		lastName:  "Kewalramani1",
		iceCream:  []string{"A", "B"},
	}
	var y person
	y.firstName = "Naveen2"
	y.lastName = "Kewalramani2"
	y.iceCream = []string{"C", "D"}
	fmt.Println(x)
	fmt.Println(y)
	for _, v := range append(x.iceCream, y.iceCream...) {
		fmt.Println(v)
	}
}
