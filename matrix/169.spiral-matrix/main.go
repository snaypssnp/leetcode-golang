package main

// https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) (res []int) {
	x := len(matrix[0])
	y := len(matrix)
	l := x * y
	right := x - 1
	bottom := y - 1
	left := 0
	top := 1

	col := 0
	row := 0

	pos := "right"

	for len(res) < l {
		num := matrix[row][col]

		if pos == "right" {
			if col == right {
				pos = "bottom"
				right--
			} else {
				res = append(res, num)
				col++
			}
		} else if pos == "bottom" {

			if row == bottom {
				pos = "left"
				bottom--
			} else {
				res = append(res, num)
				row++
			}
		} else if pos == "left" {

			if col == left {
				pos = "top"
				left++
			} else {
				res = append(res, num)
				col--
			}
		} else if pos == "top" {
			res = append(res, num)
			if row == top {
				pos = "right"
				top++
				col++
			} else {

				row--
			}
		}
	}

	return
}
