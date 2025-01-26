package main

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func connectSticks(sticks []int) (res int) {
	h := &MinHeap{}

	heap.Init(h)

	for _, stick := range sticks {
		heap.Push(h, stick)
	}

	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		c := a + b

		heap.Push(h, c)

		res += c
	}

	return
}
