package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func insertionSortList(head *ListNode) *ListNode {
	top := &ListNode{}
	var prev, next, nextIter *ListNode

	current := head
	for current != nil {
		prev, next = top, top.Next

		for next != nil && current.Val >= next.Val {
			prev = next
			next = next.Next
		}
		nextIter = current.Next
		current.Next = next
		prev.Next = current

		current = nextIter
	}

	return top.Next
}
