package main

// https://leetcode.com/problems/linked-list-cycle/description/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
