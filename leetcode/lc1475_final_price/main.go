package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{9, 9, 9, 9, 9, 9}))
	fmt.Println(finalPrices([]int{9, 10, 11, 12, 13, 14}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
}

func finalPrices(prices []int) []int {
	m := len(prices)
	monotonic := make([]int, m)
	response := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		if len(monotonic) == 0 {
			response[i] = prices[i]
		} else {
			if monotonic[len(monotonic)-1] < prices[i] {
				discount := monotonic[len(monotonic)-1]
				response[i] = prices[i] - discount
			} else {
				discount := -1
				for len(monotonic) != 0 {
					if monotonic[len(monotonic)-1] > prices[i] {
						monotonic = monotonic[0 : len(monotonic)-1]
					} else {
						discount = monotonic[len(monotonic)-1]
						break
					}
				}
				if discount == -1 {
					response[i] = prices[i]
				} else {
					response[i] = prices[i] - discount
				}
			}
			monotonic = append(monotonic, prices[i])
		}
	}
	return response
}
