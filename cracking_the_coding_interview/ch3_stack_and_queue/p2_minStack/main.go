package main

import "fmt"

type IStack interface {
	Push(int)
	Pop() int
	Min() int
}

type Element struct {
	Value    int
	MinValue int
}
type StackImpl struct {
	Stack []Element
}

func constructor() StackImpl {
	return StackImpl{
		Stack: make([]Element, 0),
	}
}

func (s *StackImpl) Min() int {
	if len(s.Stack) == 0 {
		return -1
	}
	return s.Stack[len(s.Stack)-1].MinValue
}

func (s *StackImpl) Push(val int) {
	if len(s.Stack) == 0 {
		s.Stack = append(s.Stack, Element{val, val})
	} else {
		s.Stack = append(s.Stack, Element{val, min(val, s.Min())})
	}
}

func (s *StackImpl) Pop() int {
	if len(s.Stack) == 0 {
		return -1
	}
	element := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return element.Value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	st := constructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Min())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	st.Push(0)
	fmt.Println(st.Min())
}
