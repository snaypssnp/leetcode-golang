package main

// https://leetcode.com/problems/maximum-average-subarray-i/description/
func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var res int

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	res = sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]

		if sum > res {
			res = sum
		}
	}

	return float64(res) / float64(k)
}
