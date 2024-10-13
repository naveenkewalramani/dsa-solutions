package main

import "fmt"

// Student is user defined data type
type Student struct {
	rollNo        int
	name          string
	subject1Marks int
	subject2Marks int
}

func main() {
	var student1 = Student{1, "Naveen", 10, 20}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.subject1Marks)
}
