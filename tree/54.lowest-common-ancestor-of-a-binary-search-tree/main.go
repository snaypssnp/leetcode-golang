package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root

	for node != nil {
		if node.Val > q.Val && node.Val > p.Val {
			node = node.Left
		} else if node.Val < q.Val && node.Val < p.Val {
			node = node.Right
		} else {
			return node
		}
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
