package main

import "fmt"

func main() {
	s1 := sum([]int{2, 4, 6, 8, 10}...)
	fmt.Println(s1)
	s2 := product([]int{2, 4, 6, 8, 10}...)
	fmt.Println(s2)
	s3 := productRecursion(6)
	fmt.Println(s3)
	fmt.Println(evenOperate(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	fmt.Println(evenOperate(product, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
}

// @type[Slice]
// @return[Integer]
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// @type[Slice]
// @return[Integer]
func product(x ...int) int {
	product := 1
	for _, v := range x {
		product *= v
	}
	return product
}

// @type[Integer]
// @return[Integer]
func productRecursion(n int) int {
	if n == 1 {
		return n
	}
	s := productRecursion(n - 1)
	return n * s
}

func evenOperate(f func(arg1 ...int) int, arg2 ...int) int {
	y := []int{}
	for _, v := range arg2 {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
