package main

import "container/heap"

func maxSlidingWindow(nums []int, k int) (res []int) {
	q := hp{}

	for i, v := range nums[:k-1] {
		heap.Push(&q, pair{v, i})
	}

	for i := k - 1; i < len(nums); i++ {
		heap.Push(&q, pair{nums[i], i})
		for q[0].index <= i-k {
			heap.Pop(&q)
		}
		res = append(res, q[0].value)
	}

	return
}

type pair struct{ value, index int }

type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.value > b.value || (a.value == b.value && i < j)
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
