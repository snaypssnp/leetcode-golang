package main

// https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	var rows = len(board)
	var cols = len(board[0])
	var l = len(word)

	var backtracking func(row int, col int, index int) bool

	backtracking = func(row int, col int, index int) bool {
		if l == index {
			return true
		}

		if row >= rows || col >= cols || row < 0 || col < 0 {
			return false
		}

		if board[row][col] != word[index] {
			return false
		}

		var tmp = board[row][col]
		var next = index + 1

		board[row][col] = '0'

		if backtracking(row-1, col, next) {
			return true
		}

		if backtracking(row, col-1, next) {
			return true
		}

		if backtracking(row, col+1, next) {
			return true
		}

		if backtracking(row+1, col, next) {
			return true
		}

		board[row][col] = tmp

		return false
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] && backtracking(row, col, 0) {
				return true
			}
		}
	}

	return false
}
