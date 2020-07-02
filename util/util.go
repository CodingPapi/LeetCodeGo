package util

import (
	bt "github.com/CodingPapi/LeetCodeGo/binaryTree"
)

func buildBinaryTree(input []int) *bt.TreeNode {
	if len(input) == 0 {
		return nil
	}
	var root = &bt.TreeNode{Val: input[0]}
	layer := make([]*bt.TreeNode, 0)
	layer = append(layer, root)
	var j = 1
	for i := 0; i < len(layer); {
		if j >= len(input) {
			break
		}
		if input[j] != bt.NULL {
			layer[i].Left = &bt.TreeNode{Val: input[j]}
			layer = append(layer, layer[i].Left)
		}
		j++
		if j >= len(input) {
			break
		}
		if input[j] != bt.NULL {
			layer[i].Right = &bt.TreeNode{Val: input[j]}
			layer = append(layer, layer[i].Right)
		}
		j++
		i++
	}
	return root
}
