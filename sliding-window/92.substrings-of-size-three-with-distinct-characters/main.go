package main

// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
func countGoodSubstrings(s string) (res int) {
	for i := 2; i < len(s); i++ {
		a, b, c := s[i-2], s[i-1], s[i]

		if a != b && a != c && b != c {
			res++
		}
	}

	return
}
