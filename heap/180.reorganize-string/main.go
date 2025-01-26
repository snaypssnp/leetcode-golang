package main

import "container/heap"

type CharFrequency struct {
	frequency int
	char      rune
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

// https://leetcode.com/problems/reorganize-string/
func reorganizeString(s string) (res string) {
	arr := []rune{}
	m := map[rune]int{}

	for _, char := range s {
		m[char]++
	}

	charFrequencyHeap := CharFrequencyHeap{}

	heap.Init(&charFrequencyHeap)

	for char, frequency := range m {
		charFrequency := CharFrequency{frequency, char}

		heap.Push(&charFrequencyHeap, charFrequency)
	}

	for charFrequencyHeap.Len() > 1 {
		a := heap.Pop(&charFrequencyHeap).(CharFrequency)
		b := heap.Pop(&charFrequencyHeap).(CharFrequency)

		a.frequency--
		b.frequency--

		if a.frequency > 0 {
			heap.Push(&charFrequencyHeap, a)
		}

		if b.frequency > 0 {
			heap.Push(&charFrequencyHeap, b)
		}

		arr = append(arr, a.char, b.char)
	}

	for charFrequencyHeap.Len() > 0 {
		a := heap.Pop(&charFrequencyHeap).(CharFrequency)

		if len(arr) > 0 && arr[len(arr)-1] == a.char {
			return ""
		}

		a.frequency--

		if a.frequency > 0 {
			heap.Push(&charFrequencyHeap, a)
		}

		arr = append(arr, a.char)
	}

	return string(arr)
}
