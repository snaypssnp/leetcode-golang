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
	findFirst := func(nums []int, target int) int {
		var l = 0
		var r = len(nums) - 1

		for l <= r {
			mid := (l + r) / 2

			if nums[mid] == target {
				if mid == 0 || nums[mid-1] != target {
					return mid
				}
				r = mid - 1
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}

	findLast := func(nums []int, target int) int {
		var last = len(nums) - 1
		var l = 0
		var r = last

		for l <= r {
			mid := (l + r) / 2

			if nums[mid] == target {
				if mid == last || nums[mid+1] != target {
					return mid
				}
				l = mid + 1
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		return -1
	}

	return []int{findFirst(nums, target), findLast(nums, target)}
}
