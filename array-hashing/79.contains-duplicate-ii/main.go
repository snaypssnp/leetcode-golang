package main

// https://leetcode.com/problems/contains-duplicate-ii/description/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for i, n := range nums {
		if v, ok := m[n]; ok {
			if absDiffInt(i, v) <= k {
				return true
			}
		}

		m[n] = i
	}

	return false
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
