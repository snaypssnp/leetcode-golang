package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHead := MyMinHeap(arr)
	heap.Init(&minHead)
	heap.Pop(&minHead)
	//heap.Pop(&minHead)
	//heap.Pop(&minHead)
	//heap.Push(&minHead, 1)

	fmt.Println(minHead)

}

type MyMinHeap []int

func (h MyMinHeap) Len() int {
	return len(h)
}

func (h MyMinHeap) Less(i, j int) bool {
	fmt.Println("Less")
	return h[i] < h[j] // max heap h[i] > h[j]
}

func (h MyMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MyMinHeap) Pop() any {
	fmt.Println("111")
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
