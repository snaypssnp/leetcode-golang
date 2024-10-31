package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/description/
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea1(height))

	fmt.Println(maxArea2(height))
}

// optimal solution O(n)
func maxArea1(height []int) int {
	var res = 0

	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])

		res = max(area, res)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

// brute force solution O(n^2)
func maxArea2(height []int) int {
	var res int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])

			if area > res {
				res = area
			}
		}
	}

	return res
}
