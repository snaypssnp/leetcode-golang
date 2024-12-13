package main

// https://leetcode.com/problems/remove-outermost-parentheses/
func removeOuterParentheses(s string) string {
	res := []rune{}
	var counter int

	for _, c := range s {
		if c == ')' {
			counter--
		}

		if counter > 0 {
			res = append(res, c)
		}

		if c == '(' {
			counter++
		}
	}

	return string(res)
}
