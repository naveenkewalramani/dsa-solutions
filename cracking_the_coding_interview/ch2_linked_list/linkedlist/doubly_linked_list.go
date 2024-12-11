package main

import "fmt"

type DoublyLinkedListNode struct {
	Val  int
	Prev *DoublyLinkedListNode
	Next *DoublyLinkedListNode
}

type IDoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
}

func Constructor() IDoublyLinkedList {
	return IDoublyLinkedList{
		Head: nil, Tail: nil,
	}
}

func createNewNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Val: val, Next: nil, Prev: nil,
	}
}

func (w *IDoublyLinkedList) InsertNodeAtHead(val int) {
	node := createNewNode(val)
	if w.Head == nil && w.Tail == nil {
		w.Head = node
		w.Tail = node
		return
	}
	w.Head.Prev = node
	node.Next = w.Head
	w.Head = node
}

func (w *IDoublyLinkedList) InsertNodeAtTail(val int) {
	node := createNewNode(val)
	if w.Head == nil && w.Tail == nil {
		w.Head = node
		w.Tail = node
		return
	}
	w.Tail.Next = node
	node.Prev = w.Tail
	w.Tail = node
}

func (w *IDoublyLinkedList) IterateFromHead() {
	var response []int
	temp := w.Head
	for temp != nil {
		response = append(response, temp.Val)
		temp = temp.Next
	}
	fmt.Println(response)
}

func (w *IDoublyLinkedList) IterateFromTail() {
	var response []int
	temp := w.Tail
	for temp != nil {
		response = append(response, temp.Val)
		temp = temp.Prev
	}
	fmt.Println(response)
}
