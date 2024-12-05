package stackImpl

type Stack interface {
	Push(element int)
	Pop() int
	Top() int
	IsEmpty() bool
}

type IStack struct {
	List []int
}

func (s *IStack) Push(element int) {
	s.List = append(s.List, element)
}

func (s *IStack) Pop() int {
	if len(s.List) == 0 {
		return -1
	}
	element := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return element
}

func (s *IStack) Top() int {
	if len(s.List) == 0 {
		return -1
	}
	element := s.List[len(s.List)-1]
	return element
}

func (s *IStack) IsEmpty() bool {
	return len(s.List) == 0
}
