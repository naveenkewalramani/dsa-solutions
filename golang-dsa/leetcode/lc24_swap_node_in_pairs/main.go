package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	swapPairs(nil)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	even := head
	odd := head.Next
	for even != nil && odd != nil {
		even.Val, odd.Val = odd.Val, even.Val
		if odd.Next == nil {
			break
		}
		even = odd.Next
		if even.Next == nil {
			break
		}
		odd = even.Next
	}
	return head

}
