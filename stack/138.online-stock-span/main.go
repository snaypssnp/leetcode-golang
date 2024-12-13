package main

// https://leetcode.com/problems/online-stock-span/
type StockSpanner struct {
	stack [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		span += this.stack[len(this.stack)-1][1]

		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, []int{price, span})

	return span
}
