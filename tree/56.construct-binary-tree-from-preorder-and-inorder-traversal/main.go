package main

func main() {
	buildTree2([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

// neetcode solution
func buildTree2(preorder []int, inorder []int) *TreeNode {
	var dfs func(preorder []int, inorder []int) *TreeNode

	dfs = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		val := preorder[0]

		node := &TreeNode{Val: val}

		k := searchIndex(inorder, val)

		node.Left = dfs(preorder[1:k+1], inorder[:k])
		node.Right = dfs(preorder[k+1:], inorder[k+1:])

		return node
	}

	return dfs(preorder, inorder)
}

func searchIndex(arr []int, val int) int {
	for i, n := range arr {
		if val == n {
			return i
		}
	}

	return -1
}

// my bad solution
func buildTree1(preorder []int, inorder []int) *TreeNode {
	m := map[int]int{}

	for i, n := range inorder {
		m[n] = i
	}

	var dfs func(queue []int) *TreeNode

	dfs = func(queue []int) *TreeNode {
		if len(queue) == 0 {
			return nil
		}

		val := queue[0]
		queue = queue[1:]

		node := &TreeNode{Val: val}

		leftQueue := []int{}
		rightQueue := []int{}

		for _, v := range queue {
			if m[v] < m[val] {
				leftQueue = append(leftQueue, v)
			} else {
				rightQueue = append(rightQueue, v)
			}
		}

		if len(leftQueue) > 0 {
			node.Left = dfs(leftQueue)
		}

		if len(rightQueue) > 0 {
			node.Right = dfs(rightQueue)
		}

		return node
	}

	return dfs(preorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
