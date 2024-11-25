package main

// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
