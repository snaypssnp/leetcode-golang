package main

// https://leetcode.com/problems/diameter-of-binary-tree/description/
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		ans = max(ans, left+right)

		return 1 + max(left, right)
	}

	height(root)

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
