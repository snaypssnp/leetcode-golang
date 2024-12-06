package main

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	stack := []rune{}

	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}
