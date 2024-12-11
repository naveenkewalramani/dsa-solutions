package main

import (
	"fmt"
)

func main() {
	singlyLinkedList()
	doublyLinkedList()
	ContainerOperations()
}

func singlyLinkedList() {
	w := SinglyLinkedListImpl{}
	w.LinkNodeSinglyLinkedList([]*SinglyLinkedListNode{
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
	fmt.Println(TraverseLinkedListRecur(w.Head, []int{}))
}

func doublyLinkedList() {
	w := IDoublyLinkedList{}
	w.InsertNodeAtHead(3)
	w.InsertNodeAtHead(2)
	w.InsertNodeAtHead(1)
	w.InsertNodeAtTail(4)
	w.InsertNodeAtTail(5)
	w.InsertNodeAtTail(6)
	w.IterateFromHead()
	w.IterateFromTail()
}
