package linkList

import (
	. "github.com/CodingPapi/LeetCodeGo/util"
)

// 83
func RemoveDuplicateFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	for head.Next != nil {
		if head.Next.Val == head.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return temp
}

// 82
func RemoveDuplicateFromSortedListII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumyHead := &ListNode{Val: -9999}
	dumyHead.Next = head
	head = dumyHead
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			dumyHead.Val = head.Next.Val
			head.Next = head.Next.Next.Next
		} else if head.Next.Val == dumyHead.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	if head.Next != nil && head.Next.Val == dumyHead.Val {
		head.Next = head.Next.Next
	}
	return dumyHead.Next
}

// 206
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := ReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return result
}

// 92
func ReverseListBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	i := 1
	backHead := head
	premHead := head
	prenHead := head

	former := head
	var cur *ListNode

	var nHead *ListNode
	var temp *ListNode
	for head != nil {
		if m == 1 {
			prenHead = head
		} else if i < m-1 {
		} else if i == m-1 {
			premHead = head
			prenHead = head.Next

			former = head.Next
		} else if i > m-1 && i <= n {
			temp = former.Next
			former.Next = cur
			cur = former
			former = temp
		} else if i == n && premHead != nil {
			premHead.Next = cur
			prenHead.Next = former
		}
		head = head.Next
		i++
	}

	return nil
}
