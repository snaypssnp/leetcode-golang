package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	minIndex := findMin(nums)

	left, right := 0, len(nums)-1

	if minIndex == 0 || (target >= nums[minIndex] && target <= nums[right]) {
		left = minIndex
	} else {
		right = minIndex - 1
	}

	for left <= right {
		mid := left + (right-left)/2

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

func findMin(nums []int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left] <= nums[right] {
			return left
		}

		mid := left + (right-left)/2

		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
