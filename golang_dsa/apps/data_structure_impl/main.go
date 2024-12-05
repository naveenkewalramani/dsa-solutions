package main

import (
	"container/heap"
	"data_structure_impl/min_heap"
	"fmt"

	generics "data_structure_impl/generics"
)

func main() {
	fmt.Println(generics.Addition(20, 40.2, 2002))
	fmt.Println(generics.Addition(20, 40.9))
	fmt.Println(generics.Subtraction(20, 40))
	fmt.Println(generics.Addition("Naveen", " Kewalramani", " is", " a", " good boy"))
	minHeapImpl()
}

func minHeapImpl() {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := min_heap.MinHeap{}
	minHeap = array
	heap.Init(&minHeap)
	fmt.Println(heap.Pop(&minHeap))
	fmt.Println(minHeap)
}
