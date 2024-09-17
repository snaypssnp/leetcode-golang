package main

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	var curr = dummy.Next
	var prev = dummy

	for curr != nil && curr.Next != nil {
		tmp := curr.Next
		curr.Next = tmp.Next
		tmp.Next = curr
		prev.Next = tmp

		prev = curr
		curr = curr.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
