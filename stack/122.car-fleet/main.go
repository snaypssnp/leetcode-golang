package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/car-fleet/
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]float64, n)

	for i := 0; i < n; i++ {
		cars[i] = [2]float64{float64(position[i]), float64(target-position[i]) / float64(speed[i])}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	var stack []float64

	for i := 0; i < n; i++ {
		for len(stack) > 0 && cars[i][1] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, cars[i][1])
	}

	fmt.Println(stack)

	return len(stack)
}
