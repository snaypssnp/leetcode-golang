package main

import "container/heap"

type FloatHeap []float64

func (h *FloatHeap) Len() int {
	return len(*h)
}

func (h *FloatHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *FloatHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *FloatHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *FloatHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://leetcode.com/problems/minimum-operations-to-halve-array-sum/
func halveArray(nums []int) (res int) {
	h := &FloatHeap{}

	heap.Init(h)
	var sum float64
	for _, num := range nums {
		f := float64(num)
		sum += f
		heap.Push(h, f)
	}

	minSum := sum / 2

	for minSum < sum {
		res++

		num := heap.Pop(h).(float64) / 2

		sum -= num

		heap.Push(h, num)
	}

	return res
}
