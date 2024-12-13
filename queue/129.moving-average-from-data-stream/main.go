package main

// https://leetcode.com/problems/moving-average-from-data-stream/
type MovingAverage struct {
	size  int
	sum   int
	queue []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size, 0, []int{}}
}

func (this *MovingAverage) Next(val int) (res float64) {
	this.queue = append(this.queue, val)
	this.sum += val

	for len(this.queue) > this.size {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}

	return float64(this.sum) / float64(len(this.queue))
}
