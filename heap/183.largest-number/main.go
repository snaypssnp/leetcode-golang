package main

import (
	"container/heap"
	"strconv"
)

type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	first, second := h[i]+h[j], h[j]+h[i]

	return first > second
}

func (h StringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StringHeap) Push(x any) {
	//fmt.Println(x)
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://leetcode.com/problems/largest-number/
func largestNumber(nums []int) (res string) {
	h := StringHeap{}

	heap.Init(&h)

	for _, num := range nums {
		str := strconv.Itoa(num)
		heap.Push(&h, str)
	}

	for h.Len() > 0 {
		res += heap.Pop(&h).(string)
	}

	if res[0] == '0' {
		res = "0"
	}

	return
}
