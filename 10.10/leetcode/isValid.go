package main

import (
	"fmt"
)

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	s := "()"

	valid := isValid(s)
	fmt.Println(valid)
	fmt.Println(pairs)
	fmt.Println(pairs[s[0]])
}

func isValid2(s string) bool {
	stack := []string{}

	for _, ch := range s {
		c := string(ch)
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == "(" && c == ")" || top == "[" && c == "]" || top == "{" && c == "}" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
