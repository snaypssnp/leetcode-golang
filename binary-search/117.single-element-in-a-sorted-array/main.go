package main

// https://leetcode.com/problems/single-element-in-a-sorted-array/
func singleNonDuplicate(nums []int) (ans int) {
	last := len(nums) - 1
	left, right := 0, last

	for left <= right {
		mid := left + (right-left)/2

		if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == last || nums[mid] != nums[mid+1]) {
			return nums[mid]
		}

		if isGood(nums, mid) {
			left = mid + 1
		} else {
			right = mid - 1

			ans = nums[mid]
		}
	}

	return
}

func isGood(nums []int, mid int) bool {
	if mid%2 == 0 {
		return nums[mid] == nums[mid+1]
	}

	return nums[mid] == nums[mid-1]
}
