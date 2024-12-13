package main

// https://leetcode.com/problems/remove-k-digits/
func removeKdigits(num string, k int) (res string) {
	stack := []rune{}

	if len(num) == k {
		return "0"
	}

	for _, v := range num {
		for len(stack) > 0 && v < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, v)
	}

	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
