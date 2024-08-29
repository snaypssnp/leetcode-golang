package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/string-compression
func main() {
	var chars = []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	compress(chars)

	for _, char := range chars {
		fmt.Println(string(char))
	}
}

func compress(chars []byte) int {
	var compressed = []byte{}
	var i = 0
	var j = 0

	for ; i < len(chars); i++ {
		if chars[i] != chars[j] {
			compressed = append(compressed, chars[j])
			num := i - j

			if num > 1 {
				compressed = append(compressed, []byte(strconv.Itoa(num))...)
			}

			j = i
		}
	}

	lastNum := i - j

	if lastNum > 0 {
		compressed = append(compressed, chars[j])

		if lastNum > 1 {
			compressed = append(compressed, []byte(strconv.Itoa(lastNum))...)
		}
	}

	for i, char := range compressed {
		chars[i] = char
	}

	return len(compressed)
}
