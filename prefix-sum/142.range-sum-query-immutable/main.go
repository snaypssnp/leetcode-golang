package main

// https://leetcode.com/problems/range-sum-query-immutable/
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	items := make([]int, len(nums))
	sum := 0

	for i, num := range nums {
		sum += num

		items[i] = sum
	}

	return NumArray{nums: items}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}

	return this.nums[right] - this.nums[left-1]
}
