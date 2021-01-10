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
	if m == 1 {
		return ReverseListPreN(head, n)
	}

	head.Next = ReverseListBetween(head.Next, m-1, n-1)

	return head
}

func ReverseListPreN(head *ListNode, n int) *ListNode {
	if head.Next == nil || n == 1 {
		return head
	}

	result := ReverseListPreN(head.Next, n-1)
	successor := head.Next.Next
	head.Next.Next = head
	head.Next = successor
	return result
}
