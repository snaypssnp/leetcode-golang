package main

import "fmt"

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	ll = removeNthFromEnd(ll, 2)

	for {
		fmt.Println(ll.Val)

		if ll.Next == nil {
			break
		}

		ll = ll.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	Ф
	dummy := &ListNode{Val: 0, Next: head}

	slow, fast := dummy, dummy

	for ; n > 0; n-- {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
