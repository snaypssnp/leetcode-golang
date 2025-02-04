package main

import "sort"

// https://leetcode.com/problems/how-many-apples-can-you-put-into-the-basket/
func maxNumberOfApples(weight []int) (res int) {
	sort.Ints(weight)

	var sum int

	for _, apple := range weight {
		if sum+apple > 5000 {
			break
		}
		sum += apple
		res++
	}

	return
}
