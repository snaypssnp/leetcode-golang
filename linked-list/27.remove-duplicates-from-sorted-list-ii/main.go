package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -101, Next: head}
	slow, fast := dummy, head

	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			num := fast.Val

			for fast != nil && fast.Val == num {
				fast = fast.Next
			}

			if fast == nil {
				slow.Next = fast
			}
		} else {
			slow.Next = fast
			fast = fast.Next
			slow = slow.Next
		}
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
