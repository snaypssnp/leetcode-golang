package main

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) (res []string) {
	var backtracking func(str string, left int, right int)

	backtracking = func(str string, left int, right int) {
		if len(str) == 2*n {
			res = append(res, str)
			return
		}

		if left < n {
			backtracking(str+"(", left+1, right)
		}

		if right < left {
			backtracking(str+")", left, right+1)
		}
	}

	backtracking("", 0, 0)

	return
}
