package main

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i, prefix := 0, 1; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i, postfix := len(nums)-1, 1; i >= 0; i-- {
		res[i] *= postfix

		postfix *= nums[i]
	}

	return res
}
