package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for ; n > 0; n-- {
		p2 = p2.Next
	}
	if p2 == nil {
		return head.Next
	}
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next
	return head
}
