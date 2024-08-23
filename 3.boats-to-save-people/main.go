package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/boats-to-save-people
func main() {
	people := []int{3, 2, 2, 1}

	output := numRescueBoats(people, 3)

	fmt.Println(output)
}

// O(n) solution
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	var boats int

	for j, i := len(people)-1, 0; j >= i; j-- {
		if (people[i] + people[j]) <= limit {
			i++
		}

		boats++
	}

	return boats
}
