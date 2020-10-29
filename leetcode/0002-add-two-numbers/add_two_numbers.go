package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, sum, c := new(ListNode), 0, 0
	node := head
	for l1 != nil || l2 != nil || c != 0 {
		sum = c
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
		c = sum / 10
	}
	return head.Next
}

func addTwoNumbersOld(l1 *ListNode, l2 *ListNode) *ListNode {
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
		aux.Next = &ListNode{Val: s % 10}
		aux = aux.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if c > 0 {
		aux.Next = &ListNode{Val: c}
	}
	return l.Next
}
