package linkedlist

import (
	"container/list"
	"fmt"
)

func ContainerOperations() {
	l := list.New()
	l.Init()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushFront(0)
	fmt.Println(l.Len())
	fmt.Println("Front element", l.Front().Value)
	fmt.Println("Back element", l.Back().Value)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
