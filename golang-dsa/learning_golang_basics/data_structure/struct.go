// struct and embedde struct
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	gender    string
}
type adult struct {
	details    person
	occupation string
}

type teen struct {
	details    person
	school     string
	rollNumber int
}

type seniors struct {
	person
	age int
}

func main() {
	x := person{"Naveen", "kewalramani", "Male"}
	fmt.Println(x, x.firstName, x.lastName)

	y := person{
		firstName: "Ayushi",
		lastName:  "Jaiswal",
		gender:    "Female",
	}
	fmt.Println(y, y.firstName, y.lastName)

	z := adult{
		details:    x,
		occupation: "Engineer",
	}
	fmt.Println(z, z.details, z.details.firstName, z.details.lastName, z.details.gender, z.occupation)

	a := teen{
		details:    y,
		school:     "Convent",
		rollNumber: 25,
	}
	fmt.Println(a, a.details, a.details.firstName, a.details.lastName, a.details.gender, a.school, a.rollNumber)

	b := seniors{x, 24}
	fmt.Println(b, b.firstName, b.lastName, b.gender, b.age)
}
