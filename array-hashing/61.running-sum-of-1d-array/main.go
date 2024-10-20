package main

// https://leetcode.com/problems/running-sum-of-1d-array/
func runningSum(nums []int) (res []int) {
	var sum = 0

	for _, n := range nums {
		sum += n
		res = append(res, sum)
	}

	return res
}
