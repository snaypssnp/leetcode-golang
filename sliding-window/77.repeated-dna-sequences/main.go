package main

// https://leetcode.com/problems/repeated-dna-sequences/description/
func findRepeatedDnaSequences(s string) (res []string) {
	m := map[string]int{}

	for i := 0; i < len(s)-10+1; i++ {
		sub := s[i : i+10]

		m[sub]++

		if m[sub] == 2 {
			res = append(res, sub)
		}
	}

	return
}
