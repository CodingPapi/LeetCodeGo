package linkList

import (
	// ut "github.com/CodingPapi/LeetCodeGo/util"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toString() string {
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
	return strings.Join(result[:], "")
}
