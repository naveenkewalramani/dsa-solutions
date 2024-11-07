package main

import "fmt"

func main() {
	s1 := a()
	s2, s3 := b()
	fmt.Println(s1, s2, s3)
}

func a() int {
	return 20
}

func b() (int, string) {
	return 30, "Maveen"
}
