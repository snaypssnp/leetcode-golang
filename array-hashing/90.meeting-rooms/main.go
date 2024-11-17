package main

import "sort"

// https://leetcode.com/problems/meeting-rooms/
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		prev, curr := intervals[i-1], intervals[i]

		if prev[1] > curr[0] {
			return false
		}
	}

	return true
}
