package main

import "fmt"

func main() {
	fmt.Println(validParentheses("(())"))
	fmt.Println(validParentheses("(({))}"))
	fmt.Println(validParentheses("()"))
	fmt.Println(validParentheses("()[]{}"))
	fmt.Println(validParentheses("(]"))
	fmt.Println(validParentheses("([])"))

}

func validParentheses(input string) bool {
	stack := []string{}
	for _, value := range input {
		char := fmt.Sprintf("%c", value)
		if len(stack) == 0 {
			stack = append(stack, char)
		} else if char == "{" || char == "[" || char == "(" {
			stack = append(stack, char)
		} else if char == "}" || char == "]" || char == ")" {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == "}" && top != "{") || (char == ")" && top != "(") || (char == "]" && top != "[") {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
