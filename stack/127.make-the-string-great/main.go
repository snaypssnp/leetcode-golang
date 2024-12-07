package main

// https://leetcode.com/problems/make-the-string-great/
func makeGood(s string) string {
	stack := []rune{}

	for _, v := range s {
		i := len(stack) - 1

		if len(stack) > 0 && (v-stack[i] == 32 || v-stack[i] == -32) {
			stack = stack[:i]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
