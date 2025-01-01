package main

// https://leetcode.com/problems/permutations-ii/
func permuteUnique(nums []int) (res [][]int) {
	var l = len(nums)
	var m = map[int]int{}

	for _, num := range nums {
		m[num]++
	}

	var backtracking func(curr []int)

	backtracking = func(curr []int) {
		if len(curr) == l {
			sub := make([]int, l)
			copy(sub, curr)
			res = append(res, sub)

			return
		}

		for num, count := range m {
			if count == 0 {
				continue
			}
			m[num]--
			curr = append(curr, num)

			backtracking(curr)
			curr = curr[:len(curr)-1]
			m[num]++

		}
	}

	backtracking([]int{})

	return
}
