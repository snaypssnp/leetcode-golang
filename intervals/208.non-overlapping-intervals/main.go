package main

import "sort"

// https://leetcode.com/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) (res int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]

		if prev[1] > curr[0] {
			res++

			if curr[1] < prev[1] {
				prev = curr
			}
		} else {
			prev = curr
		}
	}

	return
}
