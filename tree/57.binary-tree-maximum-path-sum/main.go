package main

func maxPathSum(root *TreeNode) int {
	var res = root.Val
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(dfs(node.Left), 0)

		right := max(dfs(node.Right), 0)

		res = max(node.Val+left+right, res)

		return node.Val + max(left, right)
	}

	dfs(root)

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
