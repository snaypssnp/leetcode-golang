package main

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/
func longestSubarray(nums []int) (ans int) {
	var k, i, j int
	var l = len(nums)

	for ; i < l; i++ {
		if nums[i] == 0 {
			k++
		}

		for k > 1 {
			if nums[j] == 0 {
				k--
			}
			j++
		}

		ans = max(ans, i-j)
	}

	return ans
}
