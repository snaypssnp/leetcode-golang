package main

import "container/heap"

// https://leetcode.com/problems/k-closest-points-to-origin/
type Point []int

func (p Point) Pow() int {
	x, y := p[0], p[1]

	return x*x + y*y
}

type PointHeap []Point

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	return h[i].Pow() > h[j].Pow()
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) (res [][]int) {
	pointHeap := PointHeap{}

	heap.Init(&pointHeap)

	for _, p := range points {
		point := Point(p)
		heap.Push(&pointHeap, point)

		if pointHeap.Len() > k {
			heap.Pop(&pointHeap)
		}
	}

	for pointHeap.Len() > 0 {
		point := heap.Pop(&pointHeap).(Point)

		res = append(res, []int{point[0], point[1]})
	}

	return
}
