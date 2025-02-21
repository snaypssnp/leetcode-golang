package main

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	minHeap := *h

	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (h *MinHeap) Less(i, j int) bool {
	minHeap := *h

	return minHeap[i] < minHeap[j]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]

	*h = old[:n]

	return x
}

// https://leetcode.com/problems/meeting-rooms-ii/

// Min Heap Solution
func minMeetingRooms1(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	minHeap := MinHeap{}

	heap.Init(&minHeap)

	for i := 0; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]

		if minHeap.Len() > 0 && start >= minHeap[0] {
			heap.Pop(&minHeap)
		}

		heap.Push(&minHeap, end)
	}

	return minHeap.Len()
}

// Greedy
func minMeetingRooms2(intervals [][]int) (res int) {
	list := [][2]int{}

	for _, interval := range intervals {
		list = append(list, [2]int{interval[0], 1}, [2]int{interval[1], -1})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i][0] == list[j][0] {
			return list[i][1] < list[j][1]
		}

		return list[i][0] < list[j][0]
	})
	var sum int
	for _, interval := range list {
		sum += interval[1]
		res = max(res, sum)
	}

	return
}
