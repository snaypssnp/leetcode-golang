package main

import "sort"

// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
func partitionArray(nums []int, k int) (res int) {
	sort.Ints(nums)
	var maxNum = nums[len(nums)-1]
	res++

	for i := len(nums) - 2; i >= 0; i-- {
		if maxNum-nums[i] > k {

			res++
			maxNum = nums[i]
		}
	}

	return
}
