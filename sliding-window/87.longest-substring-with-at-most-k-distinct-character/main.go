package main

// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
func lengthOfLongestSubstringKDistinct(s string, k int) (res int) {
	var m = map[uint8]int{}

	for i, j := 0, 0; i < len(s); i++ {
		m[s[i]]++

		if len(m) > k {
			slow := s[j]

			if m[slow] == 1 {
				delete(m, slow)
			} else {
				m[slow]--
			}
			j++
			continue
		}

		res = max(res, i-j+1)
	}

	return
}
