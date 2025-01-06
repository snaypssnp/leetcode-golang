package main

import "strings"

// https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	m1 := map[byte]string{}
	m2 := map[string]byte{}

	for i := 0; i < len(words); i++ {
		word, char := words[i], pattern[i]
		v1, ok1 := m1[char]
		v2, ok2 := m2[word]

		if (ok1 || ok2) && (v1 != word || v2 != char) {
			return false
		}

		m1[char] = word
		m2[word] = char
	}
	return true
}
