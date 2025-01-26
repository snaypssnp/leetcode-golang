package main

// https://leetcode.com/problems/shuffle-the-array/
func shuffle(nums []int, n int) (res []int) {
	l := len(nums) / 2

	for i := 0; i < l; i++ {
		n1, n2 := nums[i], nums[i+n]

		res = append(res, n1, n2)
	}

	return
}
