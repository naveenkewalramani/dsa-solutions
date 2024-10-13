package main

import (
	"fmt"
	"reflect"
	"toptal/orange"
	"toptal/stack"
)

func main() {
	fmt.Println(swapValue(10, 20))
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(compareStructWithoutSliceMap())

	var orange orange.Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)

	st := stack.IStack{List: []int{}}
	fmt.Println(st.IsEmpty())
	fmt.Println(st.Top())
	fmt.Println(st.Pop())
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	fmt.Println(st.IsEmpty())
	fmt.Println(st.Top())
	fmt.Println(st.Pop())
	fmt.Println(st.Top())
}

func swapValue(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func reverseArray(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func compareStructWithoutSliceMap() bool {
	scenarios := []struct {
		desc    string
		element int
	}{
		{
			desc:    "element 0",
			element: 0,
		}, {
			desc:    "element 0",
			element: 0,
		},
	}
	return scenarios[0] == scenarios[1]
}

func compareStructWithSlice() bool {
	scenarios := []struct {
		desc  string
		value []int
	}{
		{desc: "element 0", value: []int{1, 2}},
		{desc: "element 0", value: []int{1, 2}},
	}
	return reflect.DeepEqual(scenarios[0], scenarios[1])
}
