package main

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	var start = 0
	var end = len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
