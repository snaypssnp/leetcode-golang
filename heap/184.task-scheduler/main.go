package main

import "container/heap"

type CharFrequency struct {
	value     byte
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

// https://leetcode.com/problems/task-scheduler/
func leastInterval(tasks []byte, n int) (res int) {
	m := map[byte]int{}
	for _, char := range tasks {
		m[char]++
	}

	h := CharFrequencyHeap{}

	heap.Init(&h)

	for char, frequency := range m {
		cf := CharFrequency{char, frequency}

		heap.Push(&h, cf)
	}

	for h.Len() > 0 {
		arr := []CharFrequency{}
		for i := 0; i <= n; i++ {
			if h.Len() > 0 {
				cf := heap.Pop(&h).(CharFrequency)

				cf.frequency--

				if cf.frequency > 0 {
					arr = append(arr, cf)
				}
				res++
			} else {
				if len(arr) > 0 {
					res++
				}
			}

		}

		for _, cf := range arr {
			heap.Push(&h, cf)
		}
	}

	return res
}
