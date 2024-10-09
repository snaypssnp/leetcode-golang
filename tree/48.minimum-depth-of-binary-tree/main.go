package main

// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// dfs solution
func minDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil {
		return 1 + minDepth(node.Right)
	}

	if node.Right == nil {
		return 1 + minDepth(node.Left)
	}

	return 1 + min(minDepth(node.Left), minDepth(node.Right))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
