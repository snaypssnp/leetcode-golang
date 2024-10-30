package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]

		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{-1, -1}
}
