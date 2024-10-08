package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	response := []int{}
	response = preOrderTraversal(root, response)
	fmt.Println(response)
}

func buildTree(nums []int) *TreeNode {
	return treeBuilder(nums, 0, len(nums)-1)
}

func treeBuilder(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = treeBuilder(nums, left, mid-1)
	root.Right = treeBuilder(nums, mid+1, right)
	return root
}

func preOrderTraversal(root *TreeNode, elements []int) []int {
	if root == nil {
		return elements
	}
	elements = append(elements, root.Val)
	elements = preOrderTraversal(root.Left, elements)
	elements = preOrderTraversal(root.Right, elements)
	return elements
}
