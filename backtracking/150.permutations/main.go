package main

import "slices"

// https://leetcode.com/problems/permutations/
func permute(nums []int) (res [][]int) {
	var backtracking func(curr []int)
	l := len(nums)

	backtracking = func(curr []int) {
		if len(curr) == l {
			sub := make([]int, l)

			copy(sub, curr)

			res = append(res, sub)
		}

		for _, num := range nums {
			if !slices.Contains(curr, num) {
				curr = append(curr, num)
				backtracking(curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtracking([]int{})

	return
}
