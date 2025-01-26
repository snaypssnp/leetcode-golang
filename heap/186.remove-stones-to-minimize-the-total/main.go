package main

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
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

// https://leetcode.com/problems/remove-stones-to-minimize-the-total/
func minStoneSum(piles []int, k int) (res int) {
	h := IntHeap{}

	heap.Init(&h)

	for _, pile := range piles {
		heap.Push(&h, pile)
	}

	for i := 0; i < k; i++ {
		pile := heap.Pop(&h).(int)

		heap.Push(&h, pile-pile/2)
	}

	for _, pile := range h {
		res += pile
	}

	return
}
