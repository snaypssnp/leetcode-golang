package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/
func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}

// Time: O(N), Space: O(N)
func moveZeroes(nums []int) {
	notZeroes := []int{}

	for _, n := range nums {
		if n != 0 {
			notZeroes = append(notZeroes, n)
		}
	}

	lenNotZeroes := len(notZeroes)

	for i := range nums {
		if i < lenNotZeroes {
			nums[i] = notZeroes[i]
		} else {
			nums[i] = 0
		}
	}
}
