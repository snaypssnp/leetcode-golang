package main

// https://leetcode.com/problems/longest-continuous-increasing-subsequence/
func findLengthOfLCIS(nums []int) int {
	res := 1
	for i, sum := 1, 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			sum = 1
			continue
		}
		sum++
		res = max(res, sum)
	}

	return res
}
