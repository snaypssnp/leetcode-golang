package main

// https://leetcode.com/problems/sort-colors/ O(3N) -> O(N)
func sortColors1(nums []int) {
	var i int

	for n := 0; n <= 2; n++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == n {
				i++
			} else if nums[j] == n {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
	}
}

// O(N)
func sortColors2(nums []int) {
	start, current, end := 0, 0, len(nums)-1

	for current <= end {
		if nums[current] == 0 {
			nums[start], nums[current] = nums[current], nums[start]
			current++
			start++
		} else if nums[current] == 2 {
			nums[end], nums[current] = nums[current], nums[end]
			end--
		} else {
			current++
		}
	}
}
