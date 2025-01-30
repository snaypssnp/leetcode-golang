package main

import "sort"

// https://leetcode.com/problems/partition-labels/
// merge intervals approach
func partitionLabels1(s string) (res []int) {
	m := map[byte][2]int{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		if v, ok := m[char]; ok {
			m[char] = [2]int{v[0], i}
		} else {
			m[char] = [2]int{i, i}
		}
	}

	intervals := [][2]int{}

	for _, interval := range m {
		intervals = append(intervals, interval)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := [][2]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		l := len(mergedIntervals) - 1
		lastMergedInterval := mergedIntervals[l]
		interval := intervals[i]

		if lastMergedInterval[1] > interval[0] {
			mergedIntervals[l] = [2]int{lastMergedInterval[0], max(lastMergedInterval[1], interval[1])}
		} else {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}

	for _, interval := range mergedIntervals {
		res = append(res, interval[1]-interval[0]+1)
	}

	return
}

// greedy approach
func partitionLabels(s string) (res []int) {
	m := map[rune]int{}

	for i, char := range s {
		m[char] = i
	}

	start, end := 0, 0

	for i, char := range s {
		end = max(end, m[char])

		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return
}
