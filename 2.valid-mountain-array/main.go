package main

import "fmt"

// https://leetcode.com/problems/valid-mountain-array
func main() {
	arr := []int{1, 3, 1, 3}
	fmt.Println(validMountainArray(arr))
}

// time complexity O(N)
func validMountainArray(arr []int) bool {
	i := 1

	for i < len(arr) && arr[i] > arr[i-1] {
		i++
	}

	if i == 1 || i == len(arr) {
		return false
	}

	for i < len(arr) && arr[i] < arr[i-1] {
		i++
	}

	return i == len(arr)
}
