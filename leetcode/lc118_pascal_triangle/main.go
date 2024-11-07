package main

import "fmt"

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
	fmt.Println(generate(4))
	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	matrix := make([][]int, numRows)
	matrix[0] = []int{1}
	if numRows == 1 {
		return matrix
	}
	matrix[1] = []int{1, 1}
	if numRows == 2 {
		return matrix
	}
	for i := 0; i < numRows-2; i++ {
		newArr := []int{1}
		for j := 0; j < len(matrix[i+1])-1; j++ {
			newArr = append(newArr, matrix[i+1][j]+matrix[i+1][j+1])
		}
		newArr = append(newArr, 1)
		matrix[i+2] = newArr
	}
	return matrix
}
