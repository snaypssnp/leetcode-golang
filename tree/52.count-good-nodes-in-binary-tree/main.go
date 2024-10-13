package main

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
func goodNodes(root *TreeNode) (count int) {
	var dfs func(node *TreeNode, max int)
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}

		if node.Val >= max {
			max = node.Val

			count++
		}
		dfs(node.Right, max)
		dfs(node.Left, max)

	}

	dfs(root, root.Val)

	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
