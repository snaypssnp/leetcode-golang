package main

// https://leetcode.com/problems/binary-tree-preorder-traversal/
func preorderTraversal1(root *TreeNode) []int {
	var traverse func(node *TreeNode)
	var result []int

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)

	return result
}

// without recursion
func preorderTraversal2(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]

		stack = stack[:last]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
