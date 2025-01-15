package main

// https://leetcode.com/problems/matrix-diagonal-sum/
func diagonalSum(mat [][]int) (res int) {
	n := len(mat) - 1

	for i := 0; i <= n; i++ {
		res += mat[i][i]

		j := n - i

		if i != j {
			res += mat[i][j]
		}
	}

	return
}
