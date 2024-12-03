package main

// https://leetcode.com/problems/next-greater-element-ii/
func nextGreaterElements(nums []int) []int {
	stack := []int{}
	l := len(nums)
	ans := make([]int, l)

	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	for i := 0; i < l*2; i++ {
		currIndex := i % l
		curr := nums[currIndex]

		for len(stack) > 0 && nums[stack[len(stack)-1]] < curr {
			top := len(stack) - 1

			ans[stack[top]] = curr

			stack = stack[:top]
		}

		if i < l {
			stack = append(stack, currIndex)
		}
	}

	return ans
}
