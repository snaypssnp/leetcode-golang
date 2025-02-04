package main

import "slices"

// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
func smallestSubsequence(s string) string {
	m := map[rune]int{}

	for i, char := range s {
		m[char] = i
	}

	stack := []rune{}

	for i, char := range s {
		if slices.Contains(stack, char) {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > char && m[stack[len(stack)-1]] > i {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, char)
	}

	return string(stack)
}
