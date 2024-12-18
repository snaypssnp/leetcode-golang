package main

// https://leetcode.com/problems/contiguous-array/
func findMaxLength(nums []int) (res int) {
	var m = map[int]int{}
	var sum int
	for i, num := range nums {
		if num == 0 {
			num = -1
		}

		sum += num

		if sum == 0 {
			res = max(i+1, res)
		} else if v, ok := m[sum]; ok {
			res = max(i-v, res)
		} else {
			m[sum] = i
		}
	}

	return
}
