package main

// https://leetcode.com/problems/palindrome-partitioning/description/
func partition(s string) (res [][]string) {
	var backtracking func(curr []string, first int)

	backtracking = func(curr []string, first int) {
		if first == len(s) {
			sub := make([]string, len(curr))
			copy(sub, curr)
			res = append(res, sub)

			return
		}

		for i := first; i < len(s); i++ {
			subStr := s[first : i+1]

			if isPalindrome(subStr) {
				curr = append(curr, subStr)
				backtracking(curr, i+1)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtracking([]string{}, 0)

	return
}

func isPalindrome(str string) bool {
	for left, right := 0, len(str)-1; left <= right; left, right = left+1, right-1 {
		if str[left] != str[right] {
			return false
		}
	}

	return true
}
