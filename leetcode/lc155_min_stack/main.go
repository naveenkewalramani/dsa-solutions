package main

import "fmt"

type Element struct {
	Val int
	Min int
}

type MinStack struct {
	List []Element
}

func Constructor() MinStack {
	st := MinStack{
		List: make([]Element, 0),
	}
	return st
}

func (this *MinStack) Push(val int) {
	element := Element{Val: val}
	if len(this.List) == 0 {
		element.Min = val
	} else {
		element.Min = min(val, this.List[len(this.List)-1].Min)
	}
	this.List = append(this.List, element)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.List = this.List[:len(this.List)-1]
}

func (this *MinStack) Top() int {
	return this.List[len(this.List)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.List[len(this.List)-1].Min
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
