package main

import (
	"fmt"
)

// https://leetcode.com/problems/missing-number/
func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

// (N×(N+1))/2
// Time: O(N), Space O(1)
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return (n*(n+1))/2 - sum
}
