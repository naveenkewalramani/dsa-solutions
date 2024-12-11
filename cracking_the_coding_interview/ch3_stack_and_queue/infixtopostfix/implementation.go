package main

type IStack struct {
	List []string
}

func constructor() IStack {
	return IStack{
		List: make([]string, 0),
	}
}

func (s *IStack) Push(value string) {
	s.List = append(s.List, value)
}

func (s *IStack) Pop() string {
	element := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return element
}

func (s *IStack) Top() string {
	element := s.List[len(s.List)-1]
	return element
}

func (s *IStack) IsEmpty() bool {
	return len(s.List) == 0
}
