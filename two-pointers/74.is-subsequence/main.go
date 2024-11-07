package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
}

// s = "aaaaaa", t = "bbaaaa"
func isSubsequence(s string, t string) bool {
	for i, j := 0, -1; i < len(s); i++ {
		j = hasSymbol(t, j+1, s[i])

		if j == -1 {
			return false
		}
	}

	return true
}

func hasSymbol(str string, i int, symbol uint8) int {
	for ; i < len(str); i++ {
		if str[i] == symbol {
			return i
		}
	}

	return -1
}
