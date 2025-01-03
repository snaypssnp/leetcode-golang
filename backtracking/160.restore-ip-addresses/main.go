package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/restore-ip-addresses/
func restoreIpAddresses(s string) (res []string) {

	var backtracking func(curr []string, num string, first int)

	backtracking = func(curr []string, num string, first int) {
		if !isValid(num) {
			return
		}

		if len(curr) == 4 {
			if first < len(s) {
				return
			}

			res = append(res, strings.Join(curr, "."))

			return
		}

		for i := first; i < len(s); i++ {
			str := s[first : i+1]
			curr = append(curr, str)
			backtracking(curr, str, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]string{}, "", 0)
	return
}

func isValid(num string) bool {
	if len(num) < 2 {
		return true
	}

	if num[0] == '0' {
		return false
	}

	if d, _ := strconv.Atoi(num); d > 255 {
		return false
	}

	return true
}
