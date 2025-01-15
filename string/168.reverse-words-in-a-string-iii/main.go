package main

import "strings"

// https://leetcode.com/problems/reverse-words-in-a-string-iii/
func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}

	return strings.Join(words, " ")
}

func reverse(s string) string {
	bytes := []byte(s)

	for i, j := 0, len(bytes)-1; i <= j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}
