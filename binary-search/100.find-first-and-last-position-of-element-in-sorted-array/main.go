package main

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	return []int{findFirstPosition(nums, target), finsLastPosition(nums, target)}
}

func findFirstPosition(nums []int, target int) int {
	ans := -1

	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1

			ans = mid
		}
	}

	return ans
}

func finsLastPosition(nums []int, target int) int {
	ans := -1

	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1

			ans = mid
		}
	}

	return ans
}
