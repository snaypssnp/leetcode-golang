package main

import (
	"container/heap"
	"sort"
)

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
func minSetSize1(arr []int) (res int) {
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

func minSetSize2(arr []int) (res int) {
	m := map[int]int{}
	l := len(arr)
	half := l / 2
	for _, num := range arr {
		m[num]++
	}
	nums := make([][2]int, len(m))

	var i int
	for num, frequency := range m {
		nums[i] = [2]int{num, frequency}
		i++
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] > nums[j][1]
	})

	for i := 0; i < len(nums) && l > half; i++ {
		l -= nums[i][1]
		res++
	}

	return
}
