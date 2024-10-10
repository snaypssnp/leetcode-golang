package main

// dfs solution
func isValidBST1(root *TreeNode) bool {
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

// bfs solution
func isValidBST2(node *TreeNode) bool {
	var prev *int
	var stack []*TreeNode

	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			var n = len(stack) - 1

			node = stack[n]
			stack = stack[:n]

			if prev != nil && *prev >= node.Val {
				return false
			}

			prev = &node.Val

			node = node.Right
		}
	}

	return true

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
