package main

// https://leetcode.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) (res []int) {
	stack := []int{}

	for i, v := range nums {
		for len(stack) > 0 && stack[0] <= i-k {
			stack = stack[1:]
		}

		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

		if i >= k-1 && len(stack) > 0 {
			res = append(res, nums[stack[0]])
		}
	}

	return
}
