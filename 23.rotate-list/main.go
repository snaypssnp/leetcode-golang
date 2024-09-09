package main

import "fmt"

// https://leetcode.com/problems/rotate-list/description
func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	ll = rotateRight(ll, 2)

	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	n := 0
	curr := head

	if head == nil || head.Next == nil {
		return head
	}

	for curr != nil {
		n++
		curr = curr.Next
	}

	m := k % n

	if m == 0 {
		return head
	}

	fast, slow := head, head

	for ; m > 0; m-- {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	newHead := slow.Next
	slow.Next = nil

	fast.Next = head

	return newHead

}

type ListNode struct {
	Val  int
	Next *ListNode
}
