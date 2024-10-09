package main

// https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

// it's not my solution
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left, right := height(node.Left), height(node.Right)

	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	return 1 + max(left, right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
