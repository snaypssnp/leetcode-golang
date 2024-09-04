package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1203))
}

func reverse(x int) int {
	res := 0

	for x != 0 {
		n := x % 10
		x = x / 10

		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}

		res = res*10 + n
	}

	return res
}
