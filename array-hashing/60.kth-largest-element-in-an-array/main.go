package main

import "container/heap"

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap{}

	for _, num := range nums {
		heap.Push(&minHeap, num)

		if len(minHeap) > k {
			heap.Pop(&minHeap)
		}
	}

	return heap.Pop(&minHeap).(int)
}

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	minHeap := *h

	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (h *MinHeap) Less(i, j int) bool {
	minHeap := *h

	return minHeap[i] < minHeap[j]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]

	*h = old[:n]

	return x
}
