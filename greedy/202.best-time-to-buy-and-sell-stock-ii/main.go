package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) (res int) {

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			res += prices[i] - prices[i-1]
		}
	}

	return
}
