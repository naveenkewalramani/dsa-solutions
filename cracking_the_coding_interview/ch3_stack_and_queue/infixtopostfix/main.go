package main

import "fmt"

func main() {
	fmt.Println(infixToPostfix("(a+b)*c-d+e"))
}

func infixToPostfix(s string) string {
	answer := ""
	n := len(s)
	st := constructor()
	index := 0
	for index < n {
		str := string(s[index])
		if (str >= "A" && str <= "Z") || (str >= "a" && str <= "z") || (str >= "0" && str <= "9") {
			answer += str
		} else if str == "(" {
			st.Push(str)
		} else if str == ")" {
			for !st.IsEmpty() && st.Top() == "(" {
				answer += st.Top()
				st.Pop()
			}
			st.Pop()
		} else {
			for !st.IsEmpty() && priority(str) <= priority(st.Top()) {
				answer += st.Top()
				st.Pop()
			}
			st.Push(str)
		}
		index += 1
	}
	for !st.IsEmpty() {
		answer += st.Top()
		st.Pop()
	}
	return answer
}

func priority(s string) int {
	if s == "^" {
		return 3
	} else if s == "*" || s == "/" {
		return 2
	} else if s == "+" || s == "-" {
		return 1
	}
	return -1
}
