package main

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth1(root *TreeNode) int {
	var max int

	dfs(root, 0, &max)

	return max
}

func dfs(node *TreeNode, i int, max *int) {
	if node == nil {
		return
	}
	i++

	if i > *max {
		*max = i
	}

	dfs(node.Left, i, max)
	dfs(node.Right, i, max)
}

func maxDepth2(root *TreeNode) (max int) {
	var queue []*TreeNode

	if root == nil {
		return
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		max++

		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
