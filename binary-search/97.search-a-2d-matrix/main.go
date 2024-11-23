package main

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix1(matrix [][]int, target int) bool {
	for left, right := 0, len(matrix)-1; left <= right; {
		mid := left + (right-left)/2

		curr := binarySearch(matrix[mid], target)

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

func binarySearch(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2
		curr := nums[mid]

		if curr == target {
			return curr
		}

		if curr > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return nums[0]
}
