package main

import "fmt"

// https://leetcode.com/problems/partition-list/

func main() {
	ll := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: nil}}}}}}

	newLL := partition1(ll, 3)

	for newLL != nil {
		fmt.Println(newLL.Val)
		newLL = newLL.Next
	}
}

func partition1(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head, Val: -10100}
	slow, fast := dummy, dummy

	for fast != nil && fast.Val != x {
		fast = fast.Next
	}

	for slow.Next != nil && slow.Next.Val < x {
		slow = slow.Next
	}

	for fast != nil && fast.Next != nil {
		if fast.Next.Val < x {
			fastNext, slowNext := fast.Next, slow.Next
			slow.Next = fastNext
			fast.Next = fastNext.Next
			fastNext.Next = slowNext
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}

	return dummy.Next
}

func partition2(head *ListNode, x int) *ListNode {
	dummy1, dymmy2 := &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}
	first, second := dummy1, dymmy2

	cur := head

	for cur != nil {
		tmp := cur
		cur = cur.Next

		if tmp.Val < x {
			first.Next = tmp
			first = first.Next
		} else {
			second.Next = tmp
			second = second.Next
		}

		tmp.Next = nil
	}

	if dummy1.Next == nil {
		return dymmy2.Next
	}

	if dymmy2.Next == nil {
		return dummy1.Next
	}

	first.Next = dymmy2.Next

	return dummy1.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
