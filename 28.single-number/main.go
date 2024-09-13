package main

// https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	var res int

	for _, num := range nums {
		res ^= num
	}

	return res
}
