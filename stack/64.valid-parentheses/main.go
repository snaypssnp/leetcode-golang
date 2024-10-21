package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

var m = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		b := s[i]

		if v, ok := m[b]; ok {
			if len(stack) == 0 {
				return false
			}

			n := len(stack) - 1
			c := stack[n]

			if c != v {
				return false
			}

			stack = stack[:n]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
