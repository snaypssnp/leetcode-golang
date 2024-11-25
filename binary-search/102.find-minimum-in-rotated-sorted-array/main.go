package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left] <= nums[right] {
			return nums[left]
		}

		mid := (left + right) / 2

		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
