package main

import (
	"fmt"
	"strconv"
	"strings"
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

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(9)
	tree.Insert(8)
	//tree.Insert(6)
	//tree.Insert(9)
	//tree.Insert(10)

	//fmt.Printf("%#v\n", tree.Search(27))
	//fmt.Println(tree.DFSPostOrder())

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(tree.root)

	fmt.Println(data)
	ans := deser.deserialize(data)

	//fmt.Println(ans)

	var tree2 = &BST{root: ans}

	fmt.Println(tree2.DFSPreOrder())
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *BstNode) string {
	ans := []string{}

	var dfs func(node *BstNode)

	dfs = func(node *BstNode) {
		if node == nil {
			ans = append(ans, "#")

			return
		}

		ans = append(ans, strconv.Itoa(node.Val))

		dfs(node.left)
		dfs(node.right)
	}

	dfs(root)

	return strings.Join(ans, ",")
}

func (this *Codec) deserialize(data string) *BstNode {
	queue := strings.Split(data, ",")

	if len(queue) == 0 {
		return nil
	}

	var dfs func() *BstNode

	dfs = func() *BstNode {
		num, err := strconv.Atoi(queue[0])
		queue = queue[1:]

		if err != nil {
			return nil
		}

		node := &BstNode{Val: num}

		node.left = dfs()
		node.right = dfs()

		return node
	}

	return dfs()
}
