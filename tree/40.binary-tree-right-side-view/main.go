package main

// https://leetcode.com/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var result []int

	traverse(root, &result, 0)

	return result
}

func traverse(node *TreeNode, result *[]int, i int) {
	if node == nil {
		return
	}

	if len(*result) == i {
		*result = append(*result, node.Val)
	}

	traverse(node.Right, result, i+1)
	traverse(node.Left, result, i+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
