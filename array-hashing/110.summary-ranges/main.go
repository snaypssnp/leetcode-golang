package main

import "strconv"

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) (ans []string) {
	l := len(nums) - 1

	for i, j := 0, 0; i <= l; i++ {
		if l == i || nums[i+1]-nums[i] > 1 {
			s1 := strconv.Itoa(nums[i])

			if i == j {
				ans = append(ans, s1)
			} else {
				s2 := strconv.Itoa(nums[j])
				ans = append(ans, s2+"->"+s1)
			}
			j = i + 1
		}
	}

	return
}
