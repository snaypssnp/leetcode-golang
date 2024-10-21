package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/
func main() {
	nums := []int{0, 0, 1, 3, 4}

	moveZeroes(nums)

	fmt.Println(nums)
}

// the best way
func moveZeroes(nums []int) {
	i := 0

	for _, n := range nums {
		if n != 0 {
			nums[i] = n
			i++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

// Time: O(N), Space: O(N)
func moveZeroes2(nums []int) {
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

// brute force Time: O(N^2), Space: O(1)
func moveZeroes3(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[j-1] == 0 {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

// two pointers
func moveZeroes4(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[i] == 0 {
			if nums[j] == 0 {
				continue
			}

			nums[i], nums[j] = nums[j], nums[i]
		}

		i++
	}
}
