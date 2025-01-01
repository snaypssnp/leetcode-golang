package main

// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
func minStartValue(nums []int) int {
	var sum int
	var minValue = 100

	for _, num := range nums {
		sum += num

		minValue = min(minValue, sum)
	}

	if minValue > 0 {
		return 1
	}

	return -minValue + 1
}
