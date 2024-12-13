package main

// https://leetcode.com/problems/degree-of-an-array/
func findShortestSubArray(nums []int) (res int) {
	m := map[int][]int{}

	res = len(nums)

	maxLen := 0
	for i, n := range nums {
		m[n] = append(m[n], i)

		maxLen = max(maxLen, len(m[n]))
	}

	for _, items := range m {
		if len(items) == maxLen {
			res = min(items[len(items)-1]-items[0]+1, res)
		}
	}

	return res + 1
}
