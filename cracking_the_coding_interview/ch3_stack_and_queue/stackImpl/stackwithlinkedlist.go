package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type IStackWithLinkedList struct {
	TopNode *ListNode
}

func Constructor() IStackWithLinkedList {
	return IStackWithLinkedList{
		TopNode: nil,
	}
}

func (l *IStackWithLinkedList) Push(val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = l.TopNode
	l.TopNode = newNode
}

func (l *IStackWithLinkedList) Pop() int {
	temp := l.TopNode
	l.TopNode = l.TopNode.Next
	return temp.Val
}

func (l *IStackWithLinkedList) Top() int {
	return l.TopNode.Val
}

func (l *IStackWithLinkedList) IsEmpty() bool {
	if l.TopNode == nil {
		return true
	}
	return false
}
