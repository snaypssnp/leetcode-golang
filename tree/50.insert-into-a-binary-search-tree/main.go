package main

// log(N)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	newNode := &TreeNode{Val: val}

	if node == nil {
		return newNode
	}

	for {
		if val < node.Val {
			if node.Left == nil {
				node.Left = newNode

				break
			}

			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = newNode

				break
			}

			node = node.Right
		}
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
