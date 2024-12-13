package main

// https://leetcode.com/problems/removing-stars-from-a-string/
func removeStars(s string) string {
	stack := []rune{}

	for _, char := range s {
		if len(stack) > 0 && char == '*' {
			stack = stack[:len(stack)-1]
		}

		if char != '*' {
			stack = append(stack, char)
		}
	}

	return string(stack)
}
