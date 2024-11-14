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

//go:generate go run main.go
func main() {
	n1 := NewThreeNode(1)
	n2 := NewThreeNode(2)
	n3 := NewThreeNode(3)
	n4 := NewThreeNode(4)
	n5 := NewThreeNode(5)
	n1.Left, n1.Right = n2, n3
	n3.Left, n3.Right = n4, n5
	fmt.Println(n1)
}
