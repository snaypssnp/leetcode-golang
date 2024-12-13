package main

// https://leetcode.com/problems/number-of-recent-calls/
type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) (res int) {
	this.queue = append(this.queue, t)

	for t-3000 > this.queue[0] {
		this.queue = this.queue[1:]
	}

	return len(this.queue)
}
