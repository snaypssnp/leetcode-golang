package main

// https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	l := len(nums) - 1
	prevJump := 1
	for i := 0; i < l; i++ {
		prevJump = max(prevJump-1, nums[i])

		if prevJump == 0 {
			return false
		}
	}

	return true
}
