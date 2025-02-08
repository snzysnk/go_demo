package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewThreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

type BinarySearchTree struct {
	Root *TreeNode
}

// 插入
func (t *BinarySearchTree) insert(num int) {
	cur := t.Root
	if cur == nil {
		t.Root = NewThreeNode(num)
		return
	}
	var pre *TreeNode
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val > num {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	node := NewThreeNode(num)
	if pre.Val > num {
		pre.Left = node
	} else {
		pre.Right = node
	}
}

// 层序遍历打印
func (t *BinarySearchTree) dump() {
	tree := t.Root
	result := make([]int, 0)
	queue := list.New()
	queue.PushBack(tree)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		result = append(result, node.Val)
	}
	fmt.Println(result)
}

//go:generate go run main.go
func main() {
	binarySearchTree := &BinarySearchTree{}
	binarySearchTree.insert(8)
	binarySearchTree.insert(4)
	binarySearchTree.insert(12)
	binarySearchTree.insert(2)
	binarySearchTree.insert(6)
	binarySearchTree.insert(10)
	binarySearchTree.insert(14)
	binarySearchTree.insert(1)
	binarySearchTree.insert(3)
	binarySearchTree.insert(5)
	binarySearchTree.insert(7)
	binarySearchTree.insert(9)
	binarySearchTree.insert(11)
	binarySearchTree.insert(13)
	binarySearchTree.insert(15)
	binarySearchTree.dump()
}
