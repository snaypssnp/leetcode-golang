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

// it's not my dfs solution
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth3(root.Left), maxDepth3(root.Right))
}

// it's not my solution (iterative dfa)
func maxDepth4(root *TreeNode) (res int) {
	if root == nil {
		return
	}

	stack := []StackDepthNode{{depth: 1, node: root}}

	for len(stack) > 0 {
		n := len(stack) - 1
		stackDepth := stack[n]
		stack = stack[:n]

		res = max(res, stackDepth.depth)

		if stackDepth.node.Left != nil {
			stack = append(stack, StackDepthNode{stackDepth.node.Left, stackDepth.depth + 1})
		}

		if stackDepth.node.Right != nil {
			stack = append(stack, StackDepthNode{stackDepth.node.Right, stackDepth.depth + 1})
		}
	}

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackDepthNode struct {
	node  *TreeNode
	depth int
}
