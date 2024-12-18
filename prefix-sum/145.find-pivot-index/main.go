package main

// https://leetcode.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	var sum int
	arr := make([]int, len(nums))

	for i, num := range nums {
		sum += num
		arr[i] = sum
	}

	prev, sum := 0, arr[len(arr)-1]

	for i, num := range arr {
		if prev == sum-num {
			return i
		}

		prev = arr[i]
	}

	return -1
}
