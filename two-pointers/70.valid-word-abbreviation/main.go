package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/valid-word-abbreviation/
func main() {
	fmt.Println(validWordAbbreviation2("apple", "a2e"))
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

// Time Complexity - O(N)
// Space Complexity - O(1)
func validWordAbbreviation1(word string, abbr string) bool {
	var str string
	var j int

	for _, char := range abbr {
		if isNumber(char) {
			if str == "" && char == '0' {
				return false
			}

			str += string(char)

			continue
		}

		if len(str) > 0 {
			num, _ := strconv.Atoi(str)
			str = ""
			j += num
		}

		if j >= len(word) || char != rune(word[j]) {
			return false
		}

		j++
	}

	if str != "" {
		num, _ := strconv.Atoi(str)

		return len(word)-j == num
	}

	return j == len(word)
}

// Time Complexity - O(N)
// Space Complexity - O(1)
func validWordAbbreviation2(word string, abbr string) bool {
	wordIndex := 0
	abbrIndex := 0

	for abbrIndex < len(abbr) {
		if isNumber(abbr[abbrIndex]) {
			if abbr[abbrIndex] == '0' {
				return false
			}

			num := 0

			for abbrIndex < len(abbr) && isNumber(abbr[abbrIndex]) {
				num = num*10 + int(abbr[abbrIndex]-'0')
				abbrIndex++
			}
			wordIndex += num

			continue
		}

		if wordIndex >= len(word) || word[wordIndex] != abbr[abbrIndex] {
			return false
		}

		wordIndex++
		abbrIndex++
	}

	return wordIndex == len(word) && abbrIndex == len(abbr)
}
