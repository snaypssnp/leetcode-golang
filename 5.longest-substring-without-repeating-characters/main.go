package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func main() {
	fmt.Println(lengthOfLongestSubstring1("abba"))

	fmt.Println(lengthOfLongestSubstring2("abba"))
}

// Time: O(N), Space O(N)
func lengthOfLongestSubstring1(s string) int {
	charMap := make(map[uint8]int)
	maxLength := 0

	var j = 0
	var i = 0

	for ; i < len(s); i++ {
		char := s[i]

		if _, ok := charMap[char]; ok {
			maxLength = max(maxLength, i-j)

			j = max(charMap[char]+1, j)
		}

		charMap[char] = i
	}

	return max(maxLength, i-j)
}

/*
Brute force solution
Time: O(N^2), Space O(N)
*/
func lengthOfLongestSubstring2(s string) int {
	maxLength := 0
	for i, _ := range s {
		charMap := make(map[uint8]bool)
		currencyLength := 0
		for j := i; j < len(s); j++ {
			var char = s[j]

			if _, ok := charMap[char]; ok {
				break
			}
			currencyLength++
			charMap[char] = true

		}

		maxLength = max(maxLength, currencyLength)
	}
	return maxLength
}
