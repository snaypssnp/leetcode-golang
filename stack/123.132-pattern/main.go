package main

// https://leetcode.com/problems/132-pattern/
func find132pattern(nums []int) bool {
	stack := [][]int{}
	minCurr := nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]

		for len(stack) > 0 && n >= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && n > stack[len(stack)-1][1] {
			return true
		}

		stack = append(stack, []int{n, minCurr})
		minCurr = min(minCurr, n)
	}

	return false
}
