package main

// https://leetcode.com/problems/sort-list/description
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	ll1, ll2 := head, slow.Next
	slow.Next = nil
	left, right := sortList(ll1), sortList(ll2)

	return merge(left, right)
}

func merge(ll1, ll2 *ListNode) *ListNode {
	d := &ListNode{}
	curr := d

	for ll1 != nil && ll2 != nil {
		if ll1.Val < ll2.Val {
			curr.Next = ll1
			ll1 = ll1.Next
		} else {
			curr.Next = ll2
			ll2 = ll2.Next
		}

		curr = curr.Next
	}

	for ll1 != nil {
		curr.Next = ll1

		ll1 = ll1.Next
		curr = curr.Next
	}

	for ll2 != nil {
		curr.Next = ll2

		ll2 = ll2.Next
		curr = curr.Next
	}

	return d.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
