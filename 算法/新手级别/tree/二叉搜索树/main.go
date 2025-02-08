package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func NewThreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func (t *BinarySearchTree) search(num int) *TreeNode {
	node := t.Root
	for node != nil {
		if node.Val > num {
			node = node.Left
		} else if node.Val < num {
			node = node.Right
		} else {
			break
		}
	}
	return node
}

// NewData 构建1-15的二叉搜索树
func NewData() *BinarySearchTree {
	node1 := NewThreeNode(1)
	node2 := NewThreeNode(2)
	node3 := NewThreeNode(3)
	node4 := NewThreeNode(4)
	node5 := NewThreeNode(5)
	node6 := NewThreeNode(6)
	node7 := NewThreeNode(7)
	node8 := NewThreeNode(8)
	node9 := NewThreeNode(9)
	node10 := NewThreeNode(10)
	node11 := NewThreeNode(11)
	node12 := NewThreeNode(12)
	node13 := NewThreeNode(13)
	node14 := NewThreeNode(14)
	node15 := NewThreeNode(15)
	node8.Left, node8.Right = node4, node12
	node4.Left, node4.Right = node2, node6
	node2.Left, node2.Right = node1, node3
	node6.Left, node6.Right = node5, node7
	node12.Left, node12.Right = node10, node14
	node10.Left, node10.Right = node9, node11
	node14.Left, node14.Right = node13, node15
	return &BinarySearchTree{
		Root: node8,
	}
}

//go:generate go run main.go
func main() {
	node := NewData()
	search := node.search(7)
	if search == nil {
		fmt.Println("没有搜索到")
	} else {
		fmt.Println(search.Val)
	}
}
