package main

import "fmt"

// https://leetcode.com/problems/valid-anagram
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charMap1 = make(map[rune]int)
	var charMap2 = make(map[rune]int)

	for _, char := range s {
		if _, ok := charMap1[char]; ok {
			charMap1[char]++
		} else {
			charMap1[char] = 1
		}
	}

	for _, char := range t {
		if _, ok := charMap2[char]; ok {
			charMap2[char]++
		} else {
			charMap2[char] = 1
		}
	}

	for k, v := range charMap1 {
		if v != charMap2[k] {
			return false
		}
	}

	return true
}
