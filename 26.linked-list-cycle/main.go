package main

import "fmt"

func main() {

	ll1 := &ListNode{Val: 1, Next: nil}

	ll2 := &ListNode{Val: 2, Next: nil}

	ll3 := &ListNode{Val: 3, Next: nil}

	ll1.Next = ll2

	ll2.Next = ll3

	ll3.Next = ll1

	fmt.Println(hasCycle(ll1))
}

// https://leetcode.com/problems/linked-list-cycle/description/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
