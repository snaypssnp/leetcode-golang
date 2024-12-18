package main

// https://leetcode.com/problems/number-of-valid-subarrays/
func validSubarrays(nums []int) (res int) {
	stack := []int{}

	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > num {
			res += i - stack[len(stack)-1]

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		i := len(stack) - 1
		res += len(nums) - stack[i]

		stack = stack[:i]
	}

	return res
}
