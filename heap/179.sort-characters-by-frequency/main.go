package main

import "container/heap"

type CharFrequency struct {
	char      rune
	frequency int
}

type CharFrequencyHeap []CharFrequency

func (h CharFrequencyHeap) Len() int {
	return len(h)
}

func (h CharFrequencyHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h CharFrequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CharFrequencyHeap) Push(x any) {
	*h = append(*h, x.(CharFrequency))
}

func (h *CharFrequencyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) (res string) {
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
	}

	h := CharFrequencyHeap{}

	heap.Init(&h)

	for char, frequency := range m {
		heap.Push(&h, CharFrequency{char, frequency})
	}

	arr := make([]rune, 0, len(m))

	for len(h) > 0 {
		charFrequency := heap.Pop(&h).(CharFrequency)

		for charFrequency.frequency > 0 {
			charFrequency.frequency--
			arr = append(arr, charFrequency.char)
		}
	}

	return string(arr)
}
