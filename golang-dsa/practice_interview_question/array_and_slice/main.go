package main

import "fmt"

func main() {
	fmt.Println("variableDeclareRelatedQuestions")
	variableDeclareRelatedQuestions()
	fmt.Println("changingValuesOfSlicesRelatedQuestion")
	changingValuesOfSlicesRelatedQuestion()
	fmt.Println("copyOperationRelatedQuestions")
	copyOperationRelatedQuestions()
}

func changeValue(a []int) {
	a[0] = 10
}

func variableDeclareRelatedQuestions() {
	m := [4]int{2, 2, 2, 2}  // array, len = 4, capacity = 4, initalize
	a := make([]int, 0, 10)  // slice of len = 0, capacity = 10, output = [], declare
	b := make([]int, 10)     // slice of len = 10, capacity = 10, output = [0,0,0,0,0,0,0,0,0,0], initialize
	c := make([]int, 10, 10) // slice of len = 10, capacity = 10, output = [0,0,0,0,0,0,0,0,0,0]
	fmt.Println("Value a ", a, len(a), cap(a))
	fmt.Println("Value b ", b, len(b), cap(b))
	fmt.Println("Value c ", c, len(c), cap(c))
	fmt.Println("Value m ", m, len(m), cap(m))
}

func changingValuesOfSlicesRelatedQuestion() {
	m := [4]int{2, 2, 2, 2}
	a := make([]int, 0, 10)
	b := make([]int, 10)
	c := make([]int, 10, 10)
	// Before any operation
	fmt.Println("Value a ", a, len(a), cap(a))
	fmt.Println("Value b ", b, len(b), cap(b))
	fmt.Println("Value c ", c, len(c), cap(c))
	fmt.Println("Value m ", m, len(m), cap(m))
	// Changing the values
	// changeValue(a) -> Return error because it is declared
	// fmt.Println("Value a ", a, len(a), cap(a))
	changeValue(b)
	fmt.Println("Value b ", b, len(b), cap(b))
	changeValue(c)
	fmt.Println("Value c ", c, len(c), cap(c))
	// changeValue(m) -> Return error because we cannot pass [2]int to []int
	// fmt.Println("Value m ", m, len(m), cap(m))
	// slice go to heap as the size is not determined, while array go to stack as the size is defined
	// slice has 3 elements -> pointer, len, cap
}

func copyOperationRelatedQuestions() {
	m := [4]int{2, 2, 2, 2}
	b := make([]int, 10)

	// copy with variable declare -> no copy as memory is not assigned
	var l []int
	copy(l, m[:])
	fmt.Println("Value l", l, len(l), cap(l))

	// copy with variable initalize ->  copy as per exepectation
	k := make([]int, 10, 10)
	copy(k, m[:])
	fmt.Println("Value k", k, len(k), cap(k))
	k[0] = 4
	fmt.Println("Value k", k, len(k), cap(k))
	fmt.Println("Value m", m, len(m), cap(m))

	// copying a slice to another directly by assigning - changing slice change the underlying slice
	copyB := b
	copyB[0] = 1
	fmt.Println("Values copyB", copyB, len(copyB), cap(copyB))
	fmt.Println("Value b ", b, len(b), cap(b))

	// copying an array to another directly by assiging - does not change value of original array. Works same as copy() method of slice
	copyM := m
	copyM[0] = 1
	fmt.Println("Values copyB", copyM, len(copyM), cap(copyM))
	fmt.Println("Value m", m, len(m), cap(m))
}
