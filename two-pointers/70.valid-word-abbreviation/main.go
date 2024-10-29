package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(validWordAbbreviation("apple", "a2e"))
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func validWordAbbreviation(word string, abbr string) bool {
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
