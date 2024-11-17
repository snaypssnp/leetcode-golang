package main

// https://leetcode.com/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) (res int) {
	res = 100000000
	sum := 0

	for i, j := 0, 0; i < len(nums); i++ {
		if sum < target {
			sum += nums[i]
		}

		for sum >= target {
			res = min(res, i-j+1)
			sum -= nums[j]
			j++

			if res == 1 {
				return res
			}
		}

	}

	if res == 100000000 {
		return 0
	}

	return
}
