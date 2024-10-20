package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := []string{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			ans = append(ans, "#")

			return
		}

		ans = append(ans, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	queue := strings.Split(data, ",")

	if len(queue) == 0 {
		return nil
	}

	var dfs func() *TreeNode

	dfs = func() *TreeNode {
		num, err := strconv.Atoi(queue[0])
		queue = queue[1:]

		if err != nil {
			return nil
		}

		node := &TreeNode{Val: num}

		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
