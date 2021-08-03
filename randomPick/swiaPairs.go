package randomPick

import (
	. "github.com/CodingPapi/LeetCodeGo/util"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	result := swapPairs(head.Next.Next)
	var temp *ListNode
	temp, head.Next, head.Next.Next = head.Next, result, head
	return temp
}
