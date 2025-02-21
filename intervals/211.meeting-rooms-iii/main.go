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

type RoomHeap [][]int

func (h *RoomHeap) Len() int {
	return len(*h)
}

func (h *RoomHeap) Swap(i, j int) {
	minHeap := *h

	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (h *RoomHeap) Less(i, j int) bool {
	minHeap := *h

	if minHeap[i][1] == minHeap[j][1] {
		return minHeap[i][0] < minHeap[j][0]
	}

	return minHeap[i][1] < minHeap[j][1]
}

func (h *RoomHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *RoomHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]

	*h = old[:n]

	return x
}

// https://leetcode.com/problems/meeting-rooms-iii/
func mostBooked(n int, meetings [][]int) int {
	rh := RoomHeap{}
	mh := MinHeap{}
	m := map[int]int{}

	heap.Init(&mh)
	heap.Init(&rh)

	for i := 0; i < n; i++ {
		heap.Push(&mh, i)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		for rh.Len() > 0 && rh[0][1] <= meeting[0] {
			peak := heap.Pop(&rh).([]int)
			heap.Push(&mh, peak[0])
			if peak[1] == meeting[0] {
				break
			}
		}

		if mh.Len() > 0 {
			place := heap.Pop(&mh).(int)

			heap.Push(&rh, []int{place, meeting[1]})
			m[place]++

			continue
		}

		var peak = rh[0]

		d := peak[1] - meeting[0]

		meeting[0] += d
		meeting[1] += d

		heap.Pop(&rh)
		place := peak[0]
		heap.Push(&rh, []int{place, meeting[1]})
		m[place]++

	}

	arr := make([][2]int, len(m))

	i := 0
	for k, v := range m {
		arr[i] = [2]int{k, v}
		i++
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] > arr[j][1]
	})

	return arr[0][0]
}
