package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) (res int) {
	low := 100_000

	for _, n := range prices {
		if n > low {
			res = max(res, n-low)
		} else {
			low = min(n, low)
		}
	}

	return
}
