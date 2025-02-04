package main

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) (res int) {
	sum, res := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]

		res = max(res, sum)
	}

	return
}
