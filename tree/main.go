package main

import (
	"fmt"
)

type BstNode struct {
	Val   int
	left  *BstNode
	right *BstNode
}

type BST struct {
	root *BstNode
}

func (b *BST) Insert(val int) {
	var newNode = &BstNode{Val: val}
	if b.root == nil {
		b.root = newNode
		return
	}

	var tmp = b.root

	for {
		if val < tmp.Val {
			if tmp.left == nil {
				tmp.left = newNode
				return
			}

			tmp = tmp.left
		} else if val > tmp.Val {
			if tmp.right == nil {
				tmp.right = newNode
				return
			}

			tmp = tmp.right
		} else {
			return
		}
	}
}

func (b *BST) Search(val int) bool {
	if b.root == nil {
		return false
	}

	var tmp = b.root

	for tmp != nil {
		if val < tmp.Val {
			tmp = tmp.left
		} else if val > tmp.Val {
			tmp = tmp.right
		} else {
			return true
		}
	}

	return false
}

func (b *BST) MinVal() int {
	tmp := b.root

	if tmp == nil {
		return -1
	}

	for tmp.left != nil {
		tmp = tmp.left
	}

	return tmp.Val
}

func (b *BST) delete(val int, currentNode *BstNode) *BstNode {
	if b.root == nil {
		return nil
	}

	if val < currentNode.Val {
		currentNode.left = b.delete(val, currentNode.left)
	} else if val > currentNode.Val {
		currentNode.right = b.delete(val, currentNode.right)
	} else {
		if currentNode.left == nil && currentNode.right == nil {
			return nil
		}

		if currentNode.left == nil {
			return currentNode.right
		} else if currentNode.right == nil {
			return currentNode.left
		}

		minVal := b.MinVal()
		currentNode.Val = minVal
		currentNode.right = b.delete(minVal, currentNode.right)
	}

	return currentNode
}

func (b *BST) Delete(val int) {
	b.root = b.delete(val, b.root)
}

func (b *BST) BFS() []int {
	var result []int
	var queue []*BstNode

	if b.root != nil {
		queue = append(queue, b.root)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return result
}

func (b *BST) DFSPreOrder() []int {
	var result []int
	var traverse func(node *BstNode)

	traverse = func(node *BstNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		traverse(node.left)
		traverse(node.right)
	}

	traverse(b.root)

	return result
}

func (b *BST) DFSInOrder() []int {
	var result []int
	var traverse func(node *BstNode)

	traverse = func(node *BstNode) {
		if node == nil {
			return
		}

		traverse(node.left)
		result = append(result, node.Val)
		traverse(node.right)
	}

	traverse(b.root)

	return result
}

func (b *BST) DFSPostOrder() []int {
	var result []int
	var traverse func(node *BstNode)

	traverse = func(node *BstNode) {
		if node == nil {
			return
		}

		traverse(node.left)
		traverse(node.right)
		result = append(result, node.Val)
	}

	traverse(b.root)

	return result
}

func main() {
	var tree = &BST{}

	tree.Insert(47)
	tree.Insert(21)
	tree.Insert(76)
	tree.Insert(18)
	tree.Insert(27)
	tree.Insert(52)
	tree.Insert(82)

	fmt.Printf("%#v\n", tree.Search(27))
	fmt.Println(tree.DFSPostOrder())
}
