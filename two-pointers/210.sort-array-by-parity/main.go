package main

import "sort"

// https://leetcode.com/problems/sort-array-by-parity/
func sortArrayByParity(nums []int) []int {
	sort.Slice((nums), func(i, j int) bool {
		a, b := nums[i]%2, nums[j]%2

		if a == b {
			return nums[i] < nums[j]
		}

		return b > a
	})

	return nums
}
