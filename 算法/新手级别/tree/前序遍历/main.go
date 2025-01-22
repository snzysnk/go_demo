package main

import "fmt"

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

var result = make([]int, 0)

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	result = append(result, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

//go:generate go run main.go
func main() {
	preOrder(initTree())
	fmt.Println(result)
}
