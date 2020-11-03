package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n == 2 {
		return merge(lists[0], lists[1])
	}
	middle := n / 2
	return merge(mergeKLists(lists[0:middle]), mergeKLists(lists[middle:n]))
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}
