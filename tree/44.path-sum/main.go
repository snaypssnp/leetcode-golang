package main

// https://leetcode.com/problems/path-sum/
// time O(n), space O(h)
func hasPathSum1(root *TreeNode, targetSum int) bool {
	return traverse(root, 0, targetSum)
}

func traverse(node *TreeNode, sum int, targetSum int) bool {
	if node == nil {
		return false
	}

	sum += node.Val

	if node.Left == nil && node.Right == nil && sum == targetSum {
		return true
	}

	return traverse(node.Left, sum, targetSum) || traverse(node.Right, sum, targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
