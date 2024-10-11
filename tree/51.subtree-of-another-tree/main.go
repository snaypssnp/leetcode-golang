package main

// https://leetcode.com/problems/subtree-of-another-tree/description/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}
	var node *TreeNode

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if dfs(node, subRoot) {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}

func dfs(node *TreeNode, subNode *TreeNode) bool {
	if node == nil || subNode == nil {
		return node == nil && subNode == nil
	}

	if node.Val != subNode.Val {
		return false
	}

	return dfs(node.Left, subNode.Left) && dfs(node.Right, subNode.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
