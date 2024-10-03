package main

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var d = &ListNode{}
	var prev int
	var curr = d

	for l1 != nil || l2 != nil || prev > 0 {
		var sum = prev

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			prev = 1
			sum -= 10
		} else {
			prev = 0
		}

		curr.Next = &ListNode{Val: sum}
		curr = curr.Next

	}

	return d.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
