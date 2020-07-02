package binaryTree

import (
	"fmt"
)

var NULL = -1 << 63

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreeHello() {
	fmt.Println("Hello")
}
