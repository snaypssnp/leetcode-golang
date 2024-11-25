package main

import "fmt"

// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/
func isMajorityElement(nums []int, target int) bool {
	n := findLastPosition(nums, target) - findFirtsPosition(nums, target) + 1

	return n > len(nums)/n
}

func findFirtsPosition(nums []int, target int) int {
	ans := -1
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}

	return ans
}

func findLastPosition(nums []int, target int) int {
	ans := -1
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2

		fmt.Println(left, right)

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			ans = mid
			left = mid + 1
		}
	}

	return ans
}
