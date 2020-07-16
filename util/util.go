package util

const (
	NULL = -1 << 63
	MinInt32 = -1 << 31
	MaxInt32 = 1 << 31 -1
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func BuildBinaryTree(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	var root = &TreeNode{Val: input[0]}
	layer := make([]*TreeNode, 0)
	layer = append(layer, root)
	var j = 1
	for i := 0; i < len(layer); {
		if j >= len(input) {
			break
		}
		if input[j] != NULL {
			layer[i].Left = &TreeNode{Val: input[j]}
			layer = append(layer, layer[i].Left)
		}
		j++
		if j >= len(input) {
			break
		}
		if input[j] != NULL {
			layer[i].Right = &TreeNode{Val: input[j]}
			layer = append(layer, layer[i].Right)
		}
		j++
		i++
	}
	return root
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
