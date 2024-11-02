package main

import "fmt"

func main() {
	s1 := function1()
	s2 := function2()
	s3 := function3()
	s4 := s3()
	fmt.Println(s1, s2, s3, s4)
	fmt.Printf("%T\t%T\t%T\t%T\n", s1, s2, s3, s4)
}

func function1() int {
	x := 20
	return x
}

func function2() string {
	x := "30"
	return x
}

func function3() func() string {
	return func() string {
		return "50"
	}
}
