package main

// https://leetcode.com/problems/partition-labels/
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
