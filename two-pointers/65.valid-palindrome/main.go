package main

import "strings"

// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return 'a' <= c && c <= 'z' || '0' <= c && c <= '9'
}
