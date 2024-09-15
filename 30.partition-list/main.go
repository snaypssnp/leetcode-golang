package main

// https://leetcode.com/problems/partition-list/description/
func partition(head *ListNode, x int) *ListNode {
	dummy1, dymmy2 := &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}
	first, second := dummy1, dymmy2

	cur := head

	for cur != nil {
		tmp := cur
		cur = cur.Next

		if tmp.Val < x {
			first.Next = tmp
			first = first.Next
		} else {
			second.Next = tmp
			second = second.Next
		}

		tmp.Next = nil
	}

	if dummy1.Next == nil {
		return dymmy2.Next
	}

	if dymmy2.Next == nil {
		return dummy1.Next
	}

	first.Next = dymmy2.Next

	return dummy1.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
