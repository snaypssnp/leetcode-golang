package main

// https://leetcode.com/problems/combinations/
func combine(n int, k int) (res [][]int) {
	var backtracking func(curr []int, first int)

	backtracking = func(curr []int, first int) {
		if len(curr) == k {
			sub := make([]int, k)

			copy(sub, curr)

			res = append(res, sub)

			return
		}

		for i := first; i <= n; i++ {
			curr = append(curr, i)
			backtracking(curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, 1)

	return
}
