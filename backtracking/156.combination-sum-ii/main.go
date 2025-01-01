package main

import "sort"

// https://leetcode.com/problems/combination-sum-ii/description/

func combinationSum2(candidates []int, target int) (res [][]int) {
	var l = len(candidates)
	sort.Ints(candidates)
	var backtracking func(curr []int, sum int, first int)

	backtracking = func(curr []int, sum int, first int) {
		if sum >= target {
			if sum == target {
				sub := make([]int, len(curr))

				copy(sub, curr)

				res = append(res, sub)
			}
			return
		}
		prev := -1
		for i := first; i < l; i++ {
			n := candidates[i]
			if prev == n {
				continue
			}

			curr = append(curr, n)
			backtracking(curr, sum+n, i+1)
			curr = curr[:len(curr)-1]
			prev = n
		}
	}

	backtracking([]int{}, 0, 0)

	return
}
