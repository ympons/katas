package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return dummy.Next
}

func deleteDuplicatesClone(head *ListNode) *ListNode {
	dummy := &ListNode{}
	prev, count := dummy, 1
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			count++
		} else {
			if count == 1 {
				prev.Next = &ListNode{Val: head.Val}
				prev = prev.Next
			}
			count = 1
		}
		head = head.Next
	}
	return dummy.Next
}
