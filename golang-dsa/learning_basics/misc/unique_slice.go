package main

import "fmt"

func main() {
	unique()
	unique2()
	unique2Arr()
}

// get unique elements from array
// O(2N) complexity
// order not guranteed
func unique() {
	arr := []int{1, 2, 2, 4, 5}
	uniqueArr := []int{}
	b := map[int]bool{}
	for _, v := range arr {
		b[v] = true
	}
	for k := range b {
		uniqueArr = append(uniqueArr, k)
	}
	fmt.Println(uniqueArr, b)
}

// get unique elements from array
// O(N) order guranteed
func unique2() {
	arr := []int{1, 2, 2, 4, 5}
	uniqueArr := []int{}
	b := map[int]bool{}
	for _, v := range arr {
		if b[v] != true {
			uniqueArr = append(uniqueArr, v)
			b[v] = true
		}
	}
	fmt.Println(uniqueArr, b)
}

// get unique elements from array
// When passed two slices of ints, it will return a slice containing the int that appear in either or both slices.
// The returned slice should have no duplicates.
func unique2Arr() {
	arr := []int{2, 2, 4, 5}
	arr1 := []int{7, 8, 8, 1, 2, 4, 9, 3, 25}
	uniqueArr := []int{}
	b := map[int]bool{}
	for _, v := range append(arr, arr1...) {
		if b[v] != true {
			uniqueArr = append(uniqueArr, v)
			b[v] = true
		}
	}
	fmt.Println(uniqueArr, b)
}
