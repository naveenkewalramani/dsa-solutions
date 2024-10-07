package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "doge", "dogecoin"}))
}

func longestCommonPrefix(input []string) string {
	prefix := input[0]
	for i := 1; i < len(input); i++ {
		for !strings.HasPrefix(input[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return prefix
		}
	}
	return prefix
}
