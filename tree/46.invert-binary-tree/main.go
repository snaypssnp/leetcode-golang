package main

//https://leetcode.com/problems/invert-binary-tree/

// dfs solution
func invertTree1(root *TreeNode) *TreeNode {
	treverse(root)

	return root
}

func treverse(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left

	treverse(node.Left)
	treverse(node.Right)
}

// bfs solution
func invertTree2(root *TreeNode) *TreeNode {
	var queue []*TreeNode

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		node.Left, node.Right = node.Right, node.Left
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
