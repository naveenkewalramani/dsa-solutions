package main

type IMyQueue interface {
	Enqueue(value int)
	Dequeue() int
	Peek() int
	IsEmpty() bool
}

type MyQueue struct {
	MainStack MyStack
	TempStack MyStack
}

func queueConstructor() MyQueue {
	return MyQueue{
		MainStack: stackConstructor(),
		TempStack: stackConstructor(),
	}
}
func (this *MyQueue) Enqueue(value int) {
	this.MainStack.Push(value)
}

func (this *MyQueue) Dequeue() int {
	if this.MainStack.IsEmpty() {
		return -1
	}
	for !this.MainStack.IsEmpty() {
		element := this.MainStack.Pop()
		this.TempStack.Push(element)
	}
	element := this.TempStack.Pop()
	for !this.TempStack.IsEmpty() {
		this.MainStack.Push(this.TempStack.Pop())
	}
	return element
}

func (this *MyQueue) Peek() int {
	if this.MainStack.IsEmpty() {
		return -1
	}
	for !this.MainStack.IsEmpty() {
		element := this.MainStack.Pop()
		this.TempStack.Push(element)
	}
	element := this.TempStack.Pop()
	this.MainStack.Push(element)
	for !this.TempStack.IsEmpty() {
		this.MainStack.Push(this.TempStack.Pop())
	}
	return element
}

func (this *MyQueue) IsEmpty() bool {
	return this.MainStack.IsEmpty()
}

// stack implementation

type MyStack struct {
	List []int
}

func stackConstructor() MyStack {
	return MyStack{
		List: make([]int, 0),
	}
}

func (this *MyStack) Push(value int) {
	this.List = append(this.List, value)
}

func (this *MyStack) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	element := this.List[len(this.List)-1]
	this.List = this.List[1 : len(this.List)-1]
	return element
}

func (this *MyStack) IsEmpty() bool {
	return len(this.List) == 0
}
