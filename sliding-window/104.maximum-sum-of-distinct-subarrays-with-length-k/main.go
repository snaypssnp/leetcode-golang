package main

// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
func maximumSubarraySum(nums []int, k int) (res int64) {
	m := map[int]int{}
	var sum int

	for i := 0; i < len(nums); i++ {
		j := i - k + 1

		cur := nums[i]

		if _, ok := m[cur]; !ok {
			sum += cur
		}

		m[cur]++

		if i < k-1 {
			continue
		}

		if len(m) == k {
			res = max(res, int64(sum))
		}

		prev := nums[j]

		if _, ok := m[prev]; ok {
			m[prev]--

			if m[prev] == 0 {
				delete(m, prev)
				sum -= prev
			}
		}
	}

	return
}
