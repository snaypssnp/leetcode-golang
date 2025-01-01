package main

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) (res [][]int) {
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

		for j := first; j < len(candidates); j++ {
			num := candidates[j]

			curr = append(curr, num)

			backtracking(curr, sum+num, j)

			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, 0, 0)

	return
}
