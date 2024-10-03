package main

import "fmt"

// https://leetcode.com/problems/palindrome-linked-list/description/
func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}

	fmt.Println(isPalindrome(ll))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode

	for slow != nil {
		temp := slow
		slow = slow.Next
		temp.Next = prev
		prev = temp
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}

		prev = prev.Next
		head = head.Next
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
