package main

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/
func main() {
	words := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(words))
}

// Time: (N * M), Space: O(1)
func longestCommonPrefix(strs []string) string {
	for i, _ := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
