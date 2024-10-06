package main

// https://leetcode.com/problems/binary-tree-right-side-view/
func rightSideView1(root *TreeNode) []int {
	var result []int

	traverse(root, &result, 0)

	return result
}

func traverse(node *TreeNode, result *[]int, i int) {
	if node == nil {
		return
	}

	if len(*result) == i {
		*result = append(*result, node.Val)
	}

	traverse(node.Right, result, i+1)
	traverse(node.Left, result, i+1)
}

func rightSideView2(root *TreeNode) (result []int) {
	var stack []*TreeNode

	if root == nil {
		return
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1].Val)

		for n := len(stack); n > 0; n-- {
			node := stack[0]

			stack = stack[1:]

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
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
