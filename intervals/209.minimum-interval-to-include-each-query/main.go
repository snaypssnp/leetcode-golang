package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap [][]int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	minHeap := *h

	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (h *MinHeap) Less(i, j int) bool {
	minHeap := *h

	if minHeap[i][0] == minHeap[j][0] {
		return minHeap[i][1] < minHeap[j][1]
	}

	return minHeap[i][0] < minHeap[j][0]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]

	*h = old[:n]

	return x
}

// https://leetcode.com/problems/minimum-interval-to-include-each-query/
func minInterval(intervals [][]int, queries []int) []int {
	res := make([]int, len(queries))
	l := len(intervals)

	items := make([][]int, len(queries))

	for i, q := range queries {
		items[i] = []int{q, i}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i, interval := range intervals {
		if interval[0] == 71 && interval[1] == 71 {
			fmt.Println(i, interval)
		}

	}

	minHeap := make(MinHeap, 0, len(intervals))

	heap.Init(&minHeap)
	i := 0

	for _, item := range items {
		num, index := item[0], item[1]

		for i < l && intervals[i][0] <= num {
			start, end := intervals[i][0], intervals[i][1]
			heap.Push(&minHeap, []int{end - start, end})
			i++
		}

		for minHeap.Len() > 0 && minHeap[0][1] < num {
			heap.Pop(&minHeap)
		}

		if minHeap.Len() > 0 {
			res[index] = minHeap[0][0] + 1
		} else {
			res[index] = -1
		}
	}

	return res
}
