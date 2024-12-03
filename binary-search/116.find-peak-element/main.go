package main

// https://leetcode.com/problems/find-peak-element/
func findPeakElement(nums []int) int {
	last := len(nums) - 1
	left, right := 0, last

	if last == 0 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[last] > nums[last-1] {
		return last
	}

	for left < right {
		mid := left + (right-left)/2

		if mid == 0 {
			left = mid + 1

			continue
		}

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
