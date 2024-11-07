package main

// https://leetcode.com/problems/rotate-string/
func rotateString(s string, goal string) bool {
	for i := range s {
		if s[i:]+s[:i] == goal {
			return true
		}
	}

	return false
}
