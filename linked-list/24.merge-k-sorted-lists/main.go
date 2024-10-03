package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/merge-k-sorted-lists/description/
func main() {
	ll1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}
	ll2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}}

	var head = mergeKLists([]*ListNode{ll1, ll2})

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for {
		i := getMinNodeIndex(lists)

		if i == math.MaxInt {
			break
		}

		curr := lists[i]
		lists[i] = lists[i].Next
		prev.Next = curr
		prev = curr
	}

	return dummy.Next
}

func getMinNodeIndex(lists []*ListNode) int {
	var minVal = math.MaxInt
	var minIndex = math.MaxInt

	for i, node := range lists {
		if node == nil {
			continue
		}

		if node.Val < minVal {
			minVal = node.Val
			minIndex = i
		}
	}

	return minIndex
}
