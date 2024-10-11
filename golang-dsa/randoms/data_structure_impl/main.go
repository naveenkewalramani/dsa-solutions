package main

import (
	"fmt"

	ll "data_structure_impl/linkedlist"
)

func main() {
	singlyLinkedList()
	doublyLinkedList()
	ll.ContainerOperations()
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
