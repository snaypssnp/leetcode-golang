package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list/description/
func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	ll = reverseList2(ll)
	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	var reversed *ListNode

	for head != nil {
		reversed = &ListNode{Val: head.Val, Next: reversed}
		head = head.Next
	}
	return reversed
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
