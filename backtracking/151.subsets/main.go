package main

// https://leetcode.com/problems/subsets/
func subsets(nums []int) (res [][]int) {
	l := len(nums)
	var backtracking func(curr []int, i int)

	backtracking = func(curr []int, i int) {
		if i > l {
			return
		}

		sub := make([]int, len(curr))

		copy(sub, curr)

		res = append(res, sub)

		for j := i; j < l; j++ {
			num := nums[j]
			curr = append(curr, num)
			backtracking(curr, j+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtracking([]int{}, 0)

	return
}
