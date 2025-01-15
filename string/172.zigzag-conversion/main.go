package main

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) (res string) {
	m := make([][]rune, numRows)

	n := len(s)
	row := 0

	for i := 0; i < n; {
		for j := 0; j < numRows && i < n; j++ {
			m[row] = append(m[row], rune(s[i]))
			row++
			i++
		}

		if row >= 2 {
			row -= 2
		} else {
			row = 0
		}

		for row > 0 && i < n {
			m[row] = append(m[row], rune(s[i]))
			row--
			i++
		}
	}

	for _, value := range m {
		res += string(value)
	}

	return
}
