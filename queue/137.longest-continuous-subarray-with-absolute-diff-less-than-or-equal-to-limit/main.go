package main

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) (res int) {
	inc := []int{}
	dec := []int{}

	for left, right := 0, 0; right < len(nums); right++ {
		n := nums[right]

		for len(inc) > 0 && inc[len(inc)-1] > n {
			inc = inc[:len(inc)-1]
		}

		for len(dec) > 0 && dec[len(dec)-1] < n {
			dec = dec[:len(dec)-1]
		}

		inc = append(inc, n)
		dec = append(dec, n)

		for dec[0]-inc[0] > limit {
			if dec[0] == nums[left] {
				dec = dec[1:]
			}

			if inc[0] == nums[left] {
				inc = inc[1:]
			}

			left++
		}

		res = max(res, right-left+1)
	}

	return
}
