package main

// dfs solution
func isValidBST(root *TreeNode) bool {
	var prev *int

	var dfs func(node *TreeNode) bool

	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if dfs(node.Left) == false {
			return false
		}

		if prev != nil && *prev >= node.Val {
			return false
		}

		prev = &node.Val

		if dfs(node.Right) == false {
			return false
		}

		return true
	}

	return dfs(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
