package main

// https://leetcode.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}

	index := 0

	if word[index] <= 90 {
		index = 1
	}

	var isUpper = word[index] <= 90

	for i := 1; i < len(word); i++ {
		char := word[i]

		if isUpper {
			if char > 90 {
				return false
			}

			continue
		}

		if char <= 90 {
			return false
		}
	}

	return true
}
