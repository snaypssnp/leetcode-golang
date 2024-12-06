package main

// https://leetcode.com/problems/backspace-string-compare/
func backspaceCompare(s string, t string) bool {
	return removeBackspace(s) == removeBackspace(t)
}

func removeBackspace(s string) string {
	stack := []rune{}

	for _, v := range s {
		if v == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}
