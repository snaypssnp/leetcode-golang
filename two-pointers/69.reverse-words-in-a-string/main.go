package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the  sky is  blue  "))
}

func reverseWords(s string) string {
	arr := []string{}

	var word string

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char != ' ' {
			word += string(char)
		}

		if word != "" && (char == ' ' || i == len(s)-1) {
			arr = append(arr, word)
			word = ""
		}
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return strings.Join(arr, " ")
}
