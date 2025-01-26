package main

import "container/heap"

type FrequencyNum struct {
	num       int
	frequency int
}

type MaxFrequencyHeap []FrequencyNum

func (h MaxFrequencyHeap) Len() int {
	return len(h)
}

func (h MaxFrequencyHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h MaxFrequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxFrequencyHeap) Push(x any) {
	*h = append(*h, x.(FrequencyNum))
}

func (h *MaxFrequencyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://leetcode.com/problems/reduce-array-size-to-the-half/
func minSetSize(arr []int) (res int) {
	l := len(arr)
	half := len(arr) / 2
	m := map[int]int{}
	for _, num := range arr {
		m[num]++
	}

	h := MaxFrequencyHeap{}
	heap.Init(&h)

	for num, frequency := range m {
		frequencyNum := FrequencyNum{num, frequency}
		heap.Push(&h, frequencyNum)
	}

	for l > half && h.Len() > 0 {
		frequencyNum := heap.Pop(&h).(FrequencyNum)

		l -= frequencyNum.frequency
		res++
	}

	return
}
