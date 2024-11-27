package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
func main() {
	fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange3([]int{5, 7, 7, 8, 8, 10}, 8))
}

// Time: O(n), Space: O(1)
func searchRange1(nums []int, target int) []int {
	var length = len(nums)

	for i, j := 0, length-1; i <= j; {
		if nums[i] == target && nums[j] == target {
			return []int{i, j}
		}

		if nums[i] != target {
			i++
		}

		if nums[j] != target {
			j--
		}

	}

	return []int{-1, -1}
}

// Time: O(log n), Space: O(1)
func searchRange2(nums []int, target int) []int {
	search := func(nums []int, target int) int {
		var i = 0
		var j = len(nums)

		for i < j {
			mid := (i + j) / 2

			if nums[mid] >= target {
				j = mid
			} else {
				i = mid + 1
			}
		}
		return i
	}

	l := search(nums, target)
	r := search(nums, target+1)

	if l == r {
		return []int{-1, -1}
	}

	return []int{l, r - 1}

}

func searchRange3(nums []int, target int) []int {
	return []int{findFirstPosition(nums, target), finsLastPosition(nums, target)}
}

func findFirstPosition(nums []int, target int) int {
	ans := -1

	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1

			ans = mid
		}
	}

	return ans
}

func finsLastPosition(nums []int, target int) int {
	ans := -1

	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1

			ans = mid
		}
	}

	return ans
}
