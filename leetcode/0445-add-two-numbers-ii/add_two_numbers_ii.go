package leetcode

import "github.com/ympons/katas/essentials"

type ListNode = essentials.ListNode

type Add struct {
	x, y int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n1, n2 := 0, 0
	for aux := l1; aux != nil; aux = aux.Next {
		n1++
	}
	for aux := l2; aux != nil; aux = aux.Next {
		n2++
	}
	stack := make([]Add, 0, n1+n2)
	for n1 > n2 {
		stack = append(stack, Add{l1.Val, 0})
		l1 = l1.Next
		n1--
	}
	for n2 > n1 {
		stack = append(stack, Add{l2.Val, 0})
		l2 = l2.Next
		n2--
	}
	for l1 != nil && l2 != nil {
		stack = append(stack, Add{l1.Val, l2.Val})
		l1 = l1.Next
		l2 = l2.Next
	}
	head, s, c := &ListNode{}, 0, 0
	for i := len(stack) - 1; i >= 0; i-- {
		s = stack[i].x + stack[i].y + c
		head.Next = &ListNode{Val: s % 10, Next: head.Next}
		c = s / 10
	}
	if c > 0 {
		head.Next = &ListNode{Val: c, Next: head.Next}
	}
	return head.Next
}
