package main

// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	var res int

	for _, n := range nums {
		m[n] = false
	}

	for _, n := range nums {
		_, ok := m[n-1]

		if !ok {
			var v int = 1
			j := n + 1
			for checked, ok := m[j]; ok; {
				if checked {
					break
				}
				m[j] = true
				v++
				j++

				checked, ok = m[j]
			}

			res = max(res, v)
		}
	}

	return res
}
