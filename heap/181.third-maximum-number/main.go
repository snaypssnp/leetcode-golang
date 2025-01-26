package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://leetcode.com/problems/third-maximum-number/
func thirdMax(nums []int) (res int) {
	m := map[int]bool{}

	for _, num := range nums {
		m[num] = true
	}

	h := IntHeap{}

	heap.Init(&h)

	for num, _ := range m {
		heap.Push(&h, num)
	}

	if h.Len() < 3 {
		return heap.Pop(&h).(int)
	}

	for i := 0; i < 3; i++ {
		res = heap.Pop(&h).(int)
	}

	return
}
