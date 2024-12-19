package main

// https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/
func maxSubArrayLen(nums []int, k int) (res int) {
	m := map[int]int{}
	var sum int

	for i, num := range nums {
		sum += num
		d := sum - k

		if d == 0 {
			res = max(res, i+1)
		} else if v, ok := m[d]; ok {
			res = max(res, i-v)
		}

		if _, ok := m[sum]; !ok {
			m[sum] = i
		}

	}

	return res
}
