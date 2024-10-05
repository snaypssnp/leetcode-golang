package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder1(root *TreeNode) [][]int {
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

// without recursion
func levelOrder2(root *TreeNode) (result [][]int) {
	var queue []*TreeNode

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		var arr []int

		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]

			arr = append(arr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, arr)
	}

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
