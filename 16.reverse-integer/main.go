package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/reverse-integer/
func main() {
	fmt.Println(reverse(120000))
}

func reverse(x int) int {
	res := 0

	for ; x != 0; x = x / 10 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + x%10
	}

	return res
}
