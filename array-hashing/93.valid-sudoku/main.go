package main

// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	f1, f2, f3 := initMap(), initMap(), initMap()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			curr := board[i][j]

			if curr == '.' {
				continue
			}

			var k = (i / 3 * 3) + (j / 3)

			if f1(k, curr) || f2(j, curr) || f3(i, curr) {
				return false
			}
		}
	}
	return true
}

func initMap() func(i int, v byte) bool {
	m := make([]map[byte]struct{}, 9)

	return func(i int, v byte) bool {
		if m[i] == nil {
			m[i] = make(map[byte]struct{})
		}

		if _, ok := m[i][v]; ok {
			return true
		}

		m[i][v] = struct{}{}

		return false
	}
}
