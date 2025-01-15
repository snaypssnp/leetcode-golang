package main

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-words/
func topKFrequent(words []string, k int) (res []string) {
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	myHeap := MyMinHeap{}

	heap.Init(&myHeap)

	for key, count := range m {
		heap.Push(&myHeap, Word{key, count})
	}

	for myHeap.Len() > 0 && len(res) < k {
		word := heap.Pop(&myHeap).(Word)

		res = append(res, word.key)
	}

	return
}

type Word struct {
	key   string
	count int
}

type MyMinHeap []Word

func (h MyMinHeap) Len() int {
	return len(h)
}

func (h MyMinHeap) Less(i, j int) bool {
	return (h[i].count > h[j].count) ||
		(h[i].count == h[j].count && h[i].key < h[j].key)
}

func (h MyMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyMinHeap) Push(x any) {
	*h = append(*h, x.(Word))
}

func (h *MyMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
