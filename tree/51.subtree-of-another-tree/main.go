package main

// https://leetcode.com/problems/subtree-of-another-tree/description/
func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
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

func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	if dfs(root, subRoot) {
		return true
	}

	return isSubtree2(root.Left, subRoot) || isSubtree2(root.Right, subRoot)
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
