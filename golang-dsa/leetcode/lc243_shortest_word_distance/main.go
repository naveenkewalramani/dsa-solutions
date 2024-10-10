package main

import "fmt"

// https://leetcode.com/problems/shortest-word-distance/description/

func main() {
	fmt.Println(shortedDistance([]string{"a", "b", "c", "a"}, "a", "c"))
}

func shortedDistance(words []string, word1, word2 string) int {
	n := len(words)
	i1, i2 := -1, -1
	minD := n
	for index, value := range words {
		if value == word1 {
			i1 = index
		}
		if value == word2 {
			i2 = index
		}
		if i1 != -1 && i2 != -1 {
			minD = min(minD, abs(i2-i1))
		}
	}
	return minD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
