package main

// https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) (res int) {
	var l = len(nums) - 1
	far, end := 0, 0

	for i := 0; i < l; i++ {
		far = max(far, i+nums[i])

		if end == i {
			res++
			end = far
		}
	}

	return
}
