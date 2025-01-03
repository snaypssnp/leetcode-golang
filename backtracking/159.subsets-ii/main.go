package main

import "sort"

// https://leetcode.com/problems/subsets-ii/
func subsetsWithDup(nums []int) (res [][]int) {
	var n = len(nums)
	sort.Ints(nums)
	var backtracking func(curr []int, first int)

	backtracking = func(curr []int, first int) {
		sub := make([]int, len(curr))

		copy(sub, curr)

		res = append(res, sub)

		for i := first; i < n; i++ {
			if i > first && nums[i] == nums[i-1] {
				continue
			}

			curr = append(curr, nums[i])

			backtracking(curr, i+1)

			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, 0)

	return
}
