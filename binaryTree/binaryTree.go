package binaryTree

import (
	ut "github.com/CodingPapi/LeetCodeGo/util"
	"github.com/sqs/goreturns/returns"
)

// 124
func maxPathSum(root *ut.TreeNode) int {
	maxPath := ut.MinInt32
	maxPathDfs(root, &maxPath)
	return maxPath
}

func maxPathDfs(root *ut.TreeNode, max *int) int {
	if root == nil {
		return ut.MinInt32
	}
	maxLeft := maxPathDfs(root.Left, max)
	maxRight := maxPathDfs(root.Right, max)
	result := ut.MaxInt(root.Val, maxLeft + root.Val)
	result = ut.MaxInt(result, maxRight + root.Val)

	*max = ut.MaxInt(*max, result)
	*max = ut.MaxInt(*max, maxLeft + maxRight + root.Val)
	return result
}

// 236
func lowestCommonAncestor(root, p, q *ut.TreeNode) *ut.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftFound := lowestCommonAncestor(root.Left, p, q)
	rightFound := lowestCommonAncestor(root.Right, p, q)
	if leftFound == nil {
		return rightFound
	}
	if rightFound == nil {
		return leftFound
	}
	return root
}

// 102
func levelOrder(root *ut.TreeNode) [][]int {
	// bfs
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)

	leafContainer := make([]*ut.TreeNode, 0)
	leafContainer = append(leafContainer, root)

	for i := 0; i < len(leafContainer); {
		layer := make([]int, 0)
		length := len(leafContainer)
		for i < length {
			layer = append(layer, leafContainer[i].Val)
			if leafContainer[i].Left != nil {
				leafContainer = append(leafContainer, leafContainer[i].Left)
			}
			if leafContainer[i].Right != nil {
				leafContainer = append(leafContainer, leafContainer[i].Right)
			}
			i++
		}
		result = append(result, layer)
	}
	return result
}

// 107
func levelOrderBottom(root *ut.TreeNode) [][]int {
	// dfs
	result := make([][]int, 0)
	dfsOrderBottom(root, 0, &result)
	reverse := make([][]int, 0)
	for i := len(result) -1; i >= 0; i-- {
		reverse = append(reverse, result[i])
	}
	return reverse
}

func dfsOrderBottom(root *ut.TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) <= level {
		*result = append(*result, make([]int,0))

	}
	(*result)[level] = append((*result)[level], root.Val)
	dfsOrderBottom(root.Left, level + 1, result)
	dfsOrderBottom(root.Right, level + 1, result)
}

// 103
func zigzagLevelOrder(root *ut.TreeNode) [][]int {
	// bfs
	type LayerNode struct {
		root *ut.TreeNode
		next *LayerNode
	}

	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)

	layerNode := &LayerNode{root: root, next: nil}

	reverse := false
	for layerNode != nil {
		layerResult := make([]int, 0)
		var tempLayer *LayerNode
		for layerNode != nil {
			first := layerNode.root.Left
			second := layerNode.root.Right
			if reverse {
				first, second = second, first
			}
			layerResult = append(layerResult, layerNode.root.Val)
			if first != nil {
				tempLayer = &LayerNode{first, tempLayer}
			}
			if second != nil {
				tempLayer = &LayerNode{second, tempLayer}
			}
			layerNode = layerNode.next
		}
		result = append(result, layerResult)
		layerNode = tempLayer
		reverse = !reverse
	}
	return result
}

// 98
func isValidBST(root *ut.TreeNode) bool {
	if root == nil {
		return true
	}

	return checkBSTLeft(root.Left,root.Val) && checkBSTRight(root.Right, root.Val)
}

func checkBSTLeft(currNode *ut.TreeNode, lastVal int, rootVal int) bool {
	if currNode == nil {
		return true
	}
	if currNode.Val >= lastVal {
		return false
	}
	if lastVal > rootVal && currNode.Val < rootVal{
		return false
	}
	leftCheck := checkBSTLeft(currNode.Left, currNode.Val, lastVal)
	rightCheck := checkBSTRight(currNode.Right, currNode.Val, lastVal)
	return leftCheck && rightCheck
}

func checkBSTRight(currNode *ut.TreeNode, lastVal int, rootVal int) bool {
	if currNode == nil {
		return true
	}
	if currNode.Val < lastVal {
		return false
	}
	if lastVal < rootVal && currNode.Val > rootVal{
		return false
	}
	leftCheck := checkBSTLeft(currNode.Left, currNode.Val, lastVal)
	rightCheck := checkBSTRight(currNode.Right, currNode.Val, lastVal)
	return leftCheck && rightCheck
}
