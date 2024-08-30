package main

import "fmt"

// https://leetcode.com/problems/squares-of-a-sorted-array
func main() {
	var nums = []int{-5, -3, -2, -1, 2, 7}

	sortedSquares(nums)

	fmt.Println(sortedSquares(nums))
}

// O(N) solution
func sortedSquares(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	i, j, k := 0, length-1, length-1

	for i <= j {
		a := nums[i] * nums[i]
		b := nums[j] * nums[j]

		if a < b {
			result[k] = b
			j--
		} else {
			result[k] = a
			i++
		}

		k--
	}

	return result
}
