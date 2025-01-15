package main

import "container/heap"

type MyMinHeap []int

func (h MyMinHeap) Len() int {
	return len(h)
}

func (h MyMinHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MyMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MyMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	minHead := MyMinHeap([]int{})
	heap.Init(&minHead)

	for _, stone := range stones {
		heap.Push(&minHead, stone)
	}

	for minHead.Len() > 1 {
		y := heap.Pop(&minHead).(int)
		x := heap.Pop(&minHead).(int)

		if y != x {
			heap.Push(&minHead, y-x)
		}
	}

	if minHead.Len() == 0 {
		return 0
	}

	v := heap.Pop(&minHead).(int)

	return v
}
