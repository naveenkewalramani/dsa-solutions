package main

import "fmt"

func main() {
	str := "abada"
	n := len(str)
	dp := make([][]bool, n)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			dp[i][i+1] = true
			l, r = i, i+1
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if str[i] == str[j] && dp[i+1][j-1] {
				dp[i][j] = true
				l, r = i, j
			}
		}
	}

	fmt.Println(str[l : r+1])
}
