package main

// https://leetcode.com/problems/symmetric-tree/
func kthSmallest(root *TreeNode, k int) int {
	var result = traverse(root, []int{})

	return result[k-1]
}

func traverse(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	result = traverse(node.Left, result)
	result = append(result, node.Val)
	result = traverse(node.Right, result)

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
