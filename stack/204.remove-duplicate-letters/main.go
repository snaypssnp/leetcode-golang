package main

import "slices"

// https://leetcode.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) (res string) {
	m := make(map[rune]int)

	for i, char := range s {
		m[char] = i
	}

	stack := []rune{}

	for i, char := range s {
		if !slices.Contains(stack, char) {
			for len(stack) > 0 && stack[len(stack)-1] > char && m[stack[len(stack)-1]] > i {
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, char)
		}
	}

	return string(stack)
}
