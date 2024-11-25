package main

// https://leetcode.com/problems/binary-search/
func search1(nums []int, target int) int {
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

// recursion solution
func search2(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	cur := nums[mid]

	if cur == target {
		return mid
	}

	if cur < target {
		return binarySearch(nums, target, mid+1, right)
	}

	return binarySearch(nums, target, left, mid-1)
}
