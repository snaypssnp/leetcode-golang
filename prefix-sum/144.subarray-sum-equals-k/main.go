package main

// https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) (res int) {
	m := map[int]int{0: 1}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		d := sum - k

		if v, ok := m[d]; ok {
			res += v
		}

		m[sum]++
	}

	return res
}
