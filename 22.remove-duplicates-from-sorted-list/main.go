package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}

	deleteDuplicates(ll)

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
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
