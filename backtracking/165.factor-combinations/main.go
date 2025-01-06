package main

// https://leetcode.com/problems/factor-combinations/
func getFactors(n int) (res [][]int) {
	var backtracking func(curr []int, sum int, first int)

	backtracking = func(curr []int, sum int, first int) {
		if sum == 1 {
			if len(curr) > 0 {
				sub := make([]int, len(curr))

				copy(sub, curr)

				res = append(res, sub)
			}

			return
		}

		for i := first; i <= sum; i++ {
			if sum%i != 0 || i == n {
				continue
			}

			curr = append(curr, i)
			backtracking(curr, sum/i, i)
			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, n, 2)

	return
}
