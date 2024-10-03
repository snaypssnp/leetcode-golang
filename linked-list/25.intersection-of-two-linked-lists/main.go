package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-linked-lists/description/
func main() {
	ll0 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}
	ll1 := &ListNode{Val: 1, Next: ll0}
	ll2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: ll0}}

	ll3 := getIntersectionNode(ll1, ll2)

	fmt.Println(ll3.Val)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	first, second := headA, headB

	for first != second {
		if first == nil {
			first = headB
		} else {
			first = first.Next
		}

		if second == nil {
			second = headA
		} else {
			second = second.Next
		}
	}

	return first
}

type ListNode struct {
	Val  int
	Next *ListNode
}
