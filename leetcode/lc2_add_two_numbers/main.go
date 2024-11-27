// https://leetcode.com/problems/add-two-numbers/description/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode1, listNode2 := buildInput()
	m := addTwoNumbers(listNode1, listNode2)
	displayList(listNode1)
	displayList(listNode2)
	displayList(m)
}

func displayList(l1 *ListNode) {
	for l1 != nil {
		fmt.Printf("%d ", l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyNode := ListNode{Val: 0}
	dummyHead := &dummyNode
	previous := dummyHead
	for l1 != nil || l2 != nil {
		num := carry
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		carry = num / 10
		newNode := ListNode{Val: num % 10}
		dummyHead.Next = &newNode
		dummyHead = dummyHead.Next
	}
	return previous.Next
}

func buildInput() (*ListNode, *ListNode) {
	nodeSet1 := []ListNode{{Val: 2}, {Val: 4}, {Val: 3}, {Val: 6}}
	nodeSet2 := []ListNode{{Val: 5}, {Val: 6}, {Val: 4}, {Val: 5}, {Val: 5}, {Val: 5}}
	listNode1 := &nodeSet1[0]
	listNode2 := &nodeSet2[0]
	temp1, temp2 := listNode1, listNode2
	for i := 1; i < len(nodeSet1); i++ {
		listNode1.Next = &nodeSet1[i]
		listNode1 = listNode1.Next
	}
	for i := 1; i < len(nodeSet2); i++ {
		listNode2.Next = &nodeSet2[i]
		listNode2 = listNode2.Next
	}
	return temp1, temp2
}
