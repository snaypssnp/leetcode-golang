package main

// https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{}

	if root == nil {
		return true
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		for i, j := 0, len(queue)-1; i <= j; i, j = i+1, j-1 {
			if getVal(queue[i]) != getVal(queue[j]) {
				return false
			}
		}

		for n := len(queue); n > 0; n-- {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				queue = append(queue, node.Left, node.Right)
			}
		}
	}

	return true
}

func getVal(node *TreeNode) int {
	if node == nil {
		return 1000
	}

	return node.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
