package main

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	var max int

	dfs(root, 0, &max)

	return max
}

func dfs(node *TreeNode, i int, max *int) {
	if node == nil {
		return
	}
	i++

	if i > *max {
		*max = i
	}

	dfs(node.Left, i, max)
	dfs(node.Right, i, max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
