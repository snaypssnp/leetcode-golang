package main

import "math"

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
func getDecimalValue(head *ListNode) int {
	var i float64
	var sum int

	curr := head

	for curr != nil {
		curr = curr.Next
		i++
	}

	for head != nil {
		i--
		sum += int(math.Pow(2, i)) * head.Val
		head = head.Next

	}

	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}
