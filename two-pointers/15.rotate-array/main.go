package main

import "fmt"

// https://leetcode.com/problems/rotate-array/
func main() {
	nums := []int{1, 2, 3, 4, 5}

	rotate(nums, 6)

	fmt.Println(nums)
}

// Time: O(N)
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
