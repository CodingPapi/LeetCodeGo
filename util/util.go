package util

import (
	"strconv"
	"strings"
)

const (
	NULL     = -1 << 63
	MinInt32 = -1 << 31
	MaxInt32 = 1<<31 - 1
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (l *ListNode) ToString() string {
	result := []string{}
	if l == nil {
		return strings.Join(result[:], "")
	}

	curNode := l.Next
	result = append(result, strconv.Itoa(l.Val))

	for curNode != nil {
		result = append(result, strconv.Itoa(curNode.Val))
		curNode = curNode.Next
	}
	return strings.Join(result[:], ",")
}

func BuildLinkList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &ListNode{Val: input[0]}
	temp := result

	for i := 1; i < len(input); i++ {
		temp.Next = &ListNode{Val: input[i]}
		temp.Next.Val = input[i]
		temp = temp.Next
	}
	return result
}

func BuildCycledLinkList(input []int, q int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	var mark *ListNode
	result := &ListNode{Val: input[0]}
	temp := result
	if q == 0 {
		mark = result
	}

	for i := 1; i < len(input); i++ {
		temp.Next = &ListNode{Val: input[i]}
		temp.Next.Val = input[i]
		if i == q {
			mark = temp.Next
		}
		temp = temp.Next
	}
	temp.Next = mark

	return result
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
