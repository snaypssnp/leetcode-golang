package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	var result = [][]int{}

	traversal(root, &result, 0)

	return result
}

func traversal(node *TreeNode, result *[][]int, i int) {
	if node == nil {
		return
	}

	if len(*result) == i {
		*result = append(*result, []int{})
	}

	(*result)[i] = append((*result)[i], node.Val)

	traversal(node.Left, result, i+1)
	traversal(node.Right, result, i+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
