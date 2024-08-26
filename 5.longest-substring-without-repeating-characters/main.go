package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
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
