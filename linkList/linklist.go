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
		return head
	}

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
	if head.Next.Next != nil {
		head.Next.Next = head
	}
	head.Next = successor
	return result
}

// 21
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var dummy *ListNode
	if l1.Val < l2.Val {
		dummy = l1
		l1 = l1.Next
	} else {
		dummy = l2
		l2 = l2.Next
	}

	temp := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			dummy.Next = l2
			break
		}
		if l2 == nil {
			dummy.Next = l1
			break
		}
		if l1.Val < l2.Val {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}

	return temp
}

func MergeTwoSortedListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoSortedListsRecursion(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoSortedListsRecursion(l1, l2.Next)
		return l2
	}
}

//86
func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	dummySmaller := &ListNode{}
	dummyLarger := &ListNode{}

	tempA := dummySmaller
	tempB := dummyLarger

	for head != nil {
		if head.Val < x {
			dummySmaller.Next = head
			dummySmaller = dummySmaller.Next
		} else {
			dummyLarger.Next = head
			dummyLarger = dummyLarger.Next
		}
		head = head.Next
	}

	dummySmaller.Next = tempB.Next
	dummyLarger.Next = nil

	return tempA.Next
}

//148
func SortListRecursion(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	tail := FindMiddleAndCut(head)

	left := SortListRecursion(head)
	right := SortListRecursion(tail)

	result := MergeTwoSortedLists(left, right)
	return result
}

func FindMiddleAndCut(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	temp := slow.Next
	slow.Next = nil
	return temp
}
