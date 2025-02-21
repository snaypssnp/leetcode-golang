package main

import "sort"

// https://leetcode.com/problems/h-index/
func hIndex(citations []int) (res int) {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i, c := range citations {
		if c < i+1 {
			return

		}

		res = i + 1
	}

	return
}
