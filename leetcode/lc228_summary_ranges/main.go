package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{1, 2, 3, 5, 6, 7, 8, 10, 11, 12, 13, 17}))
	fmt.Println(summaryRanges([]int{1, 3, 5, 7, 9, 11, 13, 15, 17}))

}

func summaryRanges(input []int) []string {
	outputArr := []string{}
	n := len(input)
	if n == 0 {
		return outputArr
	}
	start := input[0]
	for i := 1; i < n; i += 1 {
		if input[i] != input[i-1]+1 {
			if start == input[i-1] {
				outputArr = append(outputArr, fmt.Sprintf("%d", start))
			} else {
				outputArr = append(outputArr, fmt.Sprintf("%d->%d", start, input[i-1]))
			}
			start = input[i]
		}
	}
	if start == input[n-1] {
		outputArr = append(outputArr, fmt.Sprintf("%d", start))
	} else {
		outputArr = append(outputArr, fmt.Sprintf("%d->%d", start, input[n-1]))
	}
	return outputArr
}
