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

func initTree() *TreeNode {
	n1 := NewThreeNode(1)
	n2 := NewThreeNode(2)
	n3 := NewThreeNode(3)
	n4 := NewThreeNode(4)
	n5 := NewThreeNode(5)
	n6 := NewThreeNode(6)
	n7 := NewThreeNode(7)
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7
	return n1
}

//go:generate go run main.go
func main() {
	tree := initTree()
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
