package main

// https://leetcode.com/problems/range-addition/
func getModifiedArray(length int, updates [][]int) []int {
	res := make([]int, length)

	for _, opts := range updates {
		startIdx, endIdx, inc := opts[0], opts[1], opts[2]

		for i := startIdx; i <= endIdx; i++ {
			res[i] += inc
		}
	}

	return res
}
