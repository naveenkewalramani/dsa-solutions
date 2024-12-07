package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Naveen",
		lastName:  "Kewalramani",
		iceCream:  []string{"a", "b", "c"},
	}
	p2 := person{
		firstName: "Naveen1",
		lastName:  "Kewalramani2",
		iceCream:  []string{"a1", "b1", "c1"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	z := map[string]person{}
	z[p1.lastName] = p1
	z[p2.lastName] = p2
	fmt.Println(z)

	for i, v := range z {
		fmt.Println("Last Name is:", i, "and Struct Person is: ", v)
	}
}
