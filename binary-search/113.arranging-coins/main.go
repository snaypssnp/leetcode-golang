package main

import "math"

// https://leetcode.com/problems/arranging-coins/
func arrangeCoins(n int) (ans int) {
	for left, right := 0, math.MaxInt32; left <= right; {
		mid := left + (right-left)/2

		if isGood(mid, n) {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}

	return ans
}

func isGood(val, n int) bool {
	return val*(val+1)/2 <= n
}
