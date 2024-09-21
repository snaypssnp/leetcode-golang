package main

// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/description/
func swapNodes(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for i := k; i > 1; i-- {
		fast = fast.Next
	}

	x := fast

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	y := slow

	x.Val, y.Val = y.Val, x.Val

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
