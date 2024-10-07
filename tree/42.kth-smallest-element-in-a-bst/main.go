package main

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest1(root *TreeNode, k int) int {
	var result = traverse(root, []int{})

	return result[k-1]
}

func traverse(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	result = traverse(node.Left, result)
	result = append(result, node.Val)
	result = traverse(node.Right, result)

	return result
}

// without recursion
func kthSmallest2(root *TreeNode, k int) int {
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			k--

			if k == 0 {
				return root.Val
			}

			root = root.Right
		}
	}

	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
