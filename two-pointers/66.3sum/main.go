package main

import "sort"

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			x := nums[i] + nums[j] + nums[k]

			if x > 0 {
				k--
			} else if x < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				k--
				j++

				for k > j && nums[k+1] == nums[k] {
					k--
				}

				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}

	return
}
