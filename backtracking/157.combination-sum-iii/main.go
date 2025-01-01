package main

// https://leetcode.com/problems/combination-sum-iii/description/

func combinationSum3(k int, n int) (res [][]int) {
	var backtracking func(curr []int, sum int, first int)

	backtracking = func(curr []int, sum int, first int) {
		if len(curr) == k && sum == n {
			sub := make([]int, k)
			copy(sub, curr)
			res = append(res, sub)

			return
		}

		for i := first; i <= 9; i++ {
			curr = append(curr, i)
			backtracking(curr, sum+i, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, 0, 1)

	return
}
