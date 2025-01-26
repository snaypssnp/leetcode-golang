package main

import (
	"container/heap"
	"sort"
)

type NumsHeap struct {
	distance int
	nums     []int
}

func (h NumsHeap) Len() int {
	return len(h.nums)
}

func (h *NumsHeap) Less(i, j int) bool {
	a, b, d := h.nums[i], h.nums[j], h.distance
	x, y := abs(a, d), abs(b, d)

	return x < y || (x == y && a < b)
}

func (h NumsHeap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *NumsHeap) Push(x any) {
	h.nums = append((*h).nums, x.(int))
}

func (h *NumsHeap) Pop() any {
	old := (*h).nums
	n := len(old)
	x := old[n-1]
	h.nums = old[:n-1]
	return x
}

// https://leetcode.com/problems/find-k-closest-elements/
func findClosestElements(nums []int, k int, x int) (res []int) {
	h := NumsHeap{x, []int{}}
	heap.Init(&h)

	for _, num := range nums {
		heap.Push(&h, num)

		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	for i := 0; i < k; i++ {
		num := heap.Pop(&h).(int)

		res = append(res, num)

	}
	sort.Ints(res)
	return
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
