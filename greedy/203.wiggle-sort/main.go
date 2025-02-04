package main

// https://leetcode.com/problems/wiggle-sort/
func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 0 {
			if nums[i-1] < nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		} else {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
	}
}
