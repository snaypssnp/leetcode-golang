package main

// https://leetcode.com/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) (res int) {
	r := max(piles)

	for l := 1; l <= r; {
		mid := l + (r-l)/2
		sum := 0
		for _, n := range piles {
			sum += ceil(n, mid)
		}

		if sum <= h {
			r = mid - 1
			res = mid
		} else {
			l = mid + 1
		}
	}

	return res
}

func ceil(a, b int) int {
	return (a + b - 1) / b
}

func max(nums []int) (res int) {
	for _, n := range nums {
		if res < n {
			res = n
		}
	}

	return
}
