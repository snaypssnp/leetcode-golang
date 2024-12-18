package main

// https://leetcode.com/problems/validate-stack-sequences/
func validateStackSequences(pushed []int, popped []int) bool {
	var stack = []int{}
	var i int
	for _, n := range pushed {
		stack = append(stack, n)

		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return len(stack) == 0
}
