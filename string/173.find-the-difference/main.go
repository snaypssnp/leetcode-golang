package main

// https://leetcode.com/problems/find-the-difference/
func findTheDifference(s string, t string) byte {
	m := map[rune]int{}

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		m[char]--
		if v, _ := m[char]; v < 0 {
			return byte(char)
		}
	}

	return 0
}
