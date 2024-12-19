package main

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	var sum int
	var n = len(nums)
	var res = make([]int, n)
	var hasZero bool

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			if sum == 0 {
				sum = nums[i]
			} else {
				sum *= nums[i]
			}

		} else {
			hasZero = true
		}

	}

	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			res[i] = sum
		} else if hasZero != true {
			res[i] = sum / nums[i]
		}

	}

	return res
}
