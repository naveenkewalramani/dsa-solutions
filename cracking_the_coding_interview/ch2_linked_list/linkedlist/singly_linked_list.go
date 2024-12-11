package main

type SinglyLinkedListNode struct {
	Val  int
	Next *SinglyLinkedListNode
}

type SinglyLinkedListImpl struct {
	Head *SinglyLinkedListNode
}

func (w *SinglyLinkedListImpl) CreateNodeSinglyLinkedList(val int) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{Val: val, Next: nil}
}

func (w *SinglyLinkedListImpl) LinkNodeSinglyLinkedList(nodes []*SinglyLinkedListNode) {
	head := nodes[0]
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Next = nodes[i]
	}
	w.Head = head
}

func (w *SinglyLinkedListImpl) TraverseLinkedList() []int {
	var response []int
	if w.Head == nil {
		return response
	}
	temp := w.Head
	for temp != nil {
		response = append(response, temp.Val)
		temp = temp.Next
	}
	return response
}

func TraverseLinkedListRecur(head *SinglyLinkedListNode, response []int) []int {
	if head == nil {
		return response
	}
	response = append(response, head.Val)
	response = TraverseLinkedListRecur(head.Next, response)
	return response
}

func (w *SinglyLinkedListImpl) InsertValueAtPosition(val, position int) {
	node := w.CreateNodeSinglyLinkedList(val)
	if position == 0 {
		node.Next = w.Head
		w.Head = node
		return
	}
	temp := w.Head
	for i := 0; i < position-1; i += 1 {
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = node
}

func (w *SinglyLinkedListImpl) DeleteElementAtPosition(position int) {
	if position == 0 {
		w.Head = w.Head.Next
		return
	}
	temp := w.Head
	for i := 0; i < position-1; i += 1 {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
}
