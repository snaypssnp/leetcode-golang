package main

// https://leetcode.com/problems/asteroid-collision/
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, n := range asteroids {
		for len(stack) > 0 && stack[len(stack)-1] > 0 && n < 0 {
			last := len(stack) - 1
			peak := stack[last]

			if peak < abs(n) {
				stack = stack[:last]
			} else if peak == abs(n) {
				stack = stack[:last]
				n = 0
			} else {
				n = 0
			}
		}

		if n != 0 {
			stack = append(stack, n)
		}

	}

	return stack
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return 0 - a
}
