package main

import "fmt"

type User struct {
	Name
	ContactDetails
}

type ContactDetails struct {
	Phone       string
	Extension   string
	CountryCode string
}

type Name struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func (c *ContactDetails) printPhoneNumber() {
	fmt.Println(c.Extension + "-" + c.Phone)
}
func (n *Name) printFullName() {
	fmt.Println(n.FirstName + " " + n.MiddleName + " " + n.LastName)
}

func main() {
	contact1 := ContactDetails{Extension: "+91", Phone: "8062"}
	name1 := Name{FirstName: "Naveen", MiddleName: "-", LastName: "Kewalramani"}
	contact1.printPhoneNumber()
	name1.printFullName()
	user := User{Name: name1, ContactDetails: contact1}
	user.printPhoneNumber()
	user.printFullName()
}
