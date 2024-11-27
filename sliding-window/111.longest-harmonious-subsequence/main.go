package main

import "sort"

// https://leetcode.com/problems/longest-harmonious-subsequence/
func findLHS(nums []int) (ans int) {
	sort.Ints(nums)

	left, right, l := 0, 1, len(nums)-1

	for ; right <= l; right++ {
		n1, n2 := nums[left], nums[right]
		n := n2 - n1

		if n == 0 {
			continue
		}

		if n > 1 {
			for nums[left] == n1 {
				left++
			}
		} else {
			ans = max(ans, right-left+1)
		}
	}

	return
}
