package main

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) (res int) {
	res = m[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		current := m[s[i]]
		prev := m[s[i+1]]

		if current < prev {
			res -= current
		} else {
			res += current
		}
	}

	return
}
