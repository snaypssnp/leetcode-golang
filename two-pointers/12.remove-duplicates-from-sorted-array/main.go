package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func main() {
	nums := []int{1, 1, 2}

	fmt.Println(removeDuplicates(nums))

	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
