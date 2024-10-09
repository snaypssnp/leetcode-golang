package main

func invertTree(root *TreeNode) *TreeNode {
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
