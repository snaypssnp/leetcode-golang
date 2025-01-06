package main

// https://leetcode.com/problems/isomorphic-strings/
func isIsomorphic(s string, t string) bool {
	return check(s, t) && check(t, s)
}

func check(s string, t string) bool {
	m := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]

		if v, ok := m[char]; ok {
			if v != t[i] {
				return false
			}
		} else {
			m[char] = t[i]
		}
	}

	return true
}
