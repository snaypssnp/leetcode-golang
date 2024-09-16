package main

// https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}

	leftLL, rightLL := dummy1, dummy2
	var prev *ListNode
	var prevTail *ListNode

	curr := head

	var i int

	for curr != nil {
		i++
		tmp := curr
		curr = curr.Next

		if i < left {
			leftLL.Next = tmp
			leftLL = leftLL.Next

			tmp.Next = nil
		} else if i > right {
			rightLL.Next = tmp
			rightLL = rightLL.Next

			tmp.Next = nil
		} else {
			if i == left {
				prevTail = tmp
			}
			tmp.Next = prev
			prev = tmp
		}
	}

	leftLL.Next = prev

	if prevTail != nil {
		prevTail.Next = dummy2.Next
	}

	return dummy1.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
