package main

import "fmt"

// https://leetcode.com/problems/merge-two-sorted-list
func main() {
	ll1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}
	ll2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}}

	ll3 := mergeTwoLists(ll1, ll2)

	for ll3 != nil {
		fmt.Println(ll3.Val)
		ll3 = ll3.Next
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for list1 != nil || list2 != nil {
		var curr *ListNode

		if list1 == nil {
			curr = list2
			list2 = list2.Next
		} else if list2 == nil {
			curr = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			curr = list2
			list2 = list2.Next
		} else {
			curr = list1
			list1 = list1.Next
		}

		curr.Next = nil

		prev.Next = curr
		prev = prev.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
