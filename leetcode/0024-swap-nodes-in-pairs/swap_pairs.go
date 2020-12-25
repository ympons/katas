package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func swapPairs(head *ListNode) *ListNode {
	temp := &ListNode{Next: head}
	prev := temp
	var first, second *ListNode
	for prev.Next != nil && prev.Next.Next != nil {
		first = prev.Next
		second = first.Next
		first.Next = second.Next
		prev.Next = second
		prev.Next.Next = first
		prev = first
	}
	return temp.Next
}
