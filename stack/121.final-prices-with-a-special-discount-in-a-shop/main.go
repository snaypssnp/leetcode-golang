package main

// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
func finalPrices(prices []int) []int {
	stack := []int{}

	for i := 0; i < len(prices); i++ {
		curr := prices[i]

		for len(stack) > 0 && prices[stack[len(stack)-1]] >= curr {
			n := len(stack) - 1

			prices[stack[n]] = prices[stack[n]] - curr

			stack = stack[:n]
		}

		stack = append(stack, i)
	}

	return prices
}
