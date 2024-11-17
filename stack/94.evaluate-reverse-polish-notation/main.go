package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	var stack = []int{}

	for _, token := range tokens {
		n, err := strconv.Atoi(token)

		if err != nil {
			l := len(stack)
			a, b := stack[l-1], stack[l-2]
			stack = stack[:l-2]

			if token == "+" {
				n = b + a
			} else if token == "-" {
				n = b - a
			} else if token == "*" {
				n = b * a
			} else {
				n = b / a
			}
		}

		stack = append(stack, n)
	}

	return stack[0]
}
