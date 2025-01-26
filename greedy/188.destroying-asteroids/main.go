package main

import "sort"

// https://leetcode.com/problems/destroying-asteroids/
func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, asteroid := range asteroids {
		if mass < asteroid {
			return false
		}

		mass += asteroid
	}

	return true
}
