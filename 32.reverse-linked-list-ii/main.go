package main

// https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	curr := prev.Next

	for i := 0; i < right-left; i++ {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
