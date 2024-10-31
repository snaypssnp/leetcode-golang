package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j, count := 1, 1, 1

	for ; j < len(nums); j++ {
		if nums[j-1] == nums[j] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
