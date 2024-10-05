package main

// https://leetcode.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	var result []int

	return traverse(root, result)
}

func traverse(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	result = traverse(node.Left, result)
	result = traverse(node.Right, result)

	return append(result, node.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
