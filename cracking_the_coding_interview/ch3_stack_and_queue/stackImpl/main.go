package main

import (
	"fmt"
)

func main() {
	stackWithArray()
	stackWithLL()
}

func stackWithArray() {
	s := IStack{
		List: []int{},
	}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.IsEmpty())
}
func stackWithLL() {
	s := Constructor()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.IsEmpty())
}
