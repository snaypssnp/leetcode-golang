package main

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var j int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
