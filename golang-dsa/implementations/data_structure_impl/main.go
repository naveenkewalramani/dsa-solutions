package main

import (
	"fmt"

	generics "data_structure_impl/generics"
	ll "data_structure_impl/linkedlist"
	q "data_structure_impl/queue"
	st "data_structure_impl/stack"
)

func main() {
	singlyLinkedList()
	doublyLinkedList()
	ll.ContainerOperations()
	stack()
	queue()
	fmt.Println(generics.Addition(20, 40.2, 2002))
	fmt.Println(generics.Addition(20, 40.9))
	fmt.Println(generics.Subtraction(20, 40))
	fmt.Println(generics.Addition("Naveen", " Kewalramani", " is", " a", " good boy"))
}

func singlyLinkedList() {
	w := ll.SinglyLinkedListImpl{}
	w.LinkNodeSinglyLinkedList([]*ll.SinglyLinkedListNode{
		w.CreateNodeSinglyLinkedList(10),
		w.CreateNodeSinglyLinkedList(20),
		w.CreateNodeSinglyLinkedList(30),
		w.CreateNodeSinglyLinkedList(40),
	})
	fmt.Println(w.TraverseLinkedList())
	w.InsertValueAtPosition(0, 0)
	w.InsertValueAtPosition(50, 5)
	w.DeleteElementAtPosition(0)
	w.InsertValueAtPosition(60, 5)
	w.DeleteElementAtPosition(5)
	w.DeleteElementAtPosition(3)
	fmt.Println(ll.TraverseLinkedListRecur(w.Head, []int{}))
}

func doublyLinkedList() {
	w := ll.IDoublyLinkedList{}
	w.InsertNodeAtHead(3)
	w.InsertNodeAtHead(2)
	w.InsertNodeAtHead(1)
	w.InsertNodeAtTail(4)
	w.InsertNodeAtTail(5)
	w.InsertNodeAtTail(6)
	w.IterateFromHead()
	w.IterateFromTail()
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
