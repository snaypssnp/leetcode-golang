package main

// https://leetcode.com/problems/monotonic-array/
func isMonotonic(nums []int) bool {
	increased := true
	decreased := true

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			increased = false
		}

		if nums[i] > nums[i-1] {
			decreased = false
		}
	}

	return increased || decreased
}
