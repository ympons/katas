package leetcode

import "github.com/ympons/challenges/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	aux := l
	c := 0
	for l1 != nil || l2 != nil {
		s := c
		if l1 != nil {
			s += l1.Val
		}
		if l2 != nil {
			s += l2.Val
		}
		c = s / 10
		aux.Next = &ListNode{s % 10, nil}
		aux = aux.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if c > 0 {
		aux.Next = &ListNode{c, nil}
	}
	return l.Next
}
