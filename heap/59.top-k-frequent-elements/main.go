package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/top-k-frequent-elements/description/
func main() {
	var arr = []int{3, 0, 1, 0}
	res := topKFrequent2(arr, 1)

	fmt.Println(res)
}

func topKFrequent1(nums []int, k int) []int {
	m := map[int]int{}
	res := make([]int, k)

	for _, n := range nums {
		m[n]++
	}

	for i := 0; i < k; i++ {
		var maxKey int
		var maxVal int
		for key, val := range m {
			if val > maxVal {
				maxKey = key
				maxVal = val
			}
		}
		res[i] = maxKey
		delete(m, maxKey)

	}

	return res
}

func topKFrequent2(nums []int, k int) (res []int) {
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	minHeap := MinHeap([]HeapNode{})

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
		hn := HeapNode{key, value}
		heap.Push(&minHeap, hn)

		if len(minHeap) > k {
			heap.Pop(&minHeap)
		}
	}

	for i := 0; i < k; i++ {
		v := heap.Pop(&minHeap).(HeapNode).value
		res = append(res, v)
	}

	return res
}

type HeapNode struct {
	value int
	count int
}

type MinHeap []HeapNode

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	fmt.Println("Less")
	return (*h)[i].count < (*h)[j].count
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(HeapNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]

	*h = old[:n]

	return x
}
