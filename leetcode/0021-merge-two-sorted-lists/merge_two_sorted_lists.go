package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			current.Next = &ListNode{Val: l1.Val}
			current = current.Next
			l1 = l1.Next
		} else if (l1 == nil || (l2 != nil && l2.Val <= l1.Val)) {
			current.Next = &ListNode{Val: l2.Val}
			current = current.Next
			l2 = l2.Next
		}
	}
	return head.Next
}

func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursive(l2.Next, l1)
	return l2
}
