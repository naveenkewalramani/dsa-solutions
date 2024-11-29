/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
package main

import "fmt"

type MyStack struct {
	MainQueue []int
	TempQueue []int
}

func main() {
	obj := Constructor()
	obj.Push(20)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

func Constructor() MyStack {
	stack := MyStack{
		MainQueue: make([]int, 0),
		TempQueue: make([]int, 0),
	}
	return stack
}

func (this *MyStack) Push(x int) {
	this.MainQueue = append(this.MainQueue, x)
}

func (this *MyStack) Pop() int {
	for len(this.MainQueue) > 1 {
		this.TempQueue = append(this.TempQueue, this.MainQueue[0])
		this.MainQueue = this.MainQueue[1:]
	}
	element := this.MainQueue[0]
	this.MainQueue = this.TempQueue
	this.TempQueue = []int{}
	return element
}

func (this *MyStack) Top() int {
	return this.MainQueue[len(this.MainQueue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.MainQueue) == 0
}
