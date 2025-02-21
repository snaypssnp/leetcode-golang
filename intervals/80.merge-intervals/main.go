package main

import "sort"

// https://leetcode.com/problems/merge-intervals/description/
func merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		a := intervals[i][0]
		b := intervals[i][1]

		for j := i + 1; j < len(intervals) && b >= intervals[j][0]; j++ {
			if intervals[j][1] > b {
				b = intervals[j][1]
			}

			i = j
		}
		res = append(res, []int{a, b})
	}

	return res
}
