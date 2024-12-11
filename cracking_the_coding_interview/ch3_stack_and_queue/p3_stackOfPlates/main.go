package main

import "fmt"

type IStack interface {
	Push(v int) int
	Pop() (int, int)
	CurrentSetCapacity() int
}

type StackImpl struct {
	StacksSet    [][]int
	CurrentSet   int
	ThresholdSet int
}

func constructor(threshold int) StackImpl {
	return StackImpl{
		CurrentSet:   0,
		ThresholdSet: threshold,
		StacksSet:    make([][]int, threshold),
	}
}

func (s *StackImpl) Push(v int) int {
	if len(s.StacksSet[s.CurrentSet]) == s.ThresholdSet {
		s.CurrentSet = s.CurrentSet + 1

	}
	s.StacksSet[s.CurrentSet] = append(s.StacksSet[s.CurrentSet], v)
	return s.CurrentSet
}

func (s *StackImpl) CurrentSetCapacity() int {
	return s.ThresholdSet - len(s.StacksSet[s.CurrentSet])
}

func (s *StackImpl) Pop() (int, int) {
	if s.CurrentSet == 0 && len(s.StacksSet[s.CurrentSet]) == 0 {
		// empty stack
		return -1, -1
	}
	element := s.StacksSet[s.CurrentSet][len(s.StacksSet[s.CurrentSet])-1]
	setNumber := s.CurrentSet
	s.StacksSet[s.CurrentSet] = s.StacksSet[s.CurrentSet][:len(s.StacksSet[s.CurrentSet])-1]
	if len(s.StacksSet[s.CurrentSet]) == 0 {
		s.CurrentSet -= 1
	}
	return element, setNumber
}

func main() {
	st := constructor(4)
	fmt.Println(st.CurrentSetCapacity())
	fmt.Println("Element push to set", st.Push(1))
	fmt.Println("Element push to set", st.Push(2))
	fmt.Println("Element push to set", st.Push(3))
	fmt.Println("Element push to set", st.Push(4))
	fmt.Println("Element push to set", st.Push(5))
	fmt.Println("Element push to set", st.Push(6))
	fmt.Println("Current Set", st.CurrentSet)
	fmt.Println("Capacity of current Set", st.CurrentSetCapacity())
	x, y := st.Pop()
	fmt.Println("Element pop and set", x, y)
	x, y = st.Pop()
	fmt.Println("Element pop and set", x, y)
	x, y = st.Pop()
	fmt.Println("Element pop and set", x, y)
	x, y = st.Pop()
	fmt.Println("Element pop and set", x, y)
	fmt.Println("Current Set", st.CurrentSet)
	fmt.Println("Capacity of current Set", st.CurrentSetCapacity())
	fmt.Println("Element push to set", st.Push(7))
	fmt.Println("Current Set", st.CurrentSet)
	fmt.Println("Capacity of current Set", st.CurrentSetCapacity())
	x, y = st.Pop()
	fmt.Println("Element pop and set", x, y)
}
