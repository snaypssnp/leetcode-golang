package main

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	for i, j := len(nums)-1, 1; i >= 0; i-- {
		res[i] *= j

		j *= nums[i]
	}

	return res
}
