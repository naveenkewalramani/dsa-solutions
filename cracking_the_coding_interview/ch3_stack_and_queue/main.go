package main

import (
	q "ch3_stack_and_queue/queueImpl"
	st "ch3_stack_and_queue/stackImpl"
	"fmt"
)

func main() {
	stack()
	queue()
}

func stack() {
	s := st.IStack{
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

func queue() {
	q := q.IQueue{}
	q.Init()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
}
