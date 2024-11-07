package main

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
func findDisappearedNumbers(nums []int) (res []int) {
	var m = map[int]struct{}{}
	var l = len(nums)

	for _, n := range nums {
		m[n] = struct{}{}
	}

	for i := 1; i <= l; i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}
