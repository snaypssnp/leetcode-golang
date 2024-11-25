package main

// https://leetcode.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		if binarySearch(nums, target) {
			return true
		}
	}

	return false
}

func binarySearch(nums []int, target int) bool {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2
		curr := nums[mid]

		if curr == target {
			return true
		}

		if curr > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
