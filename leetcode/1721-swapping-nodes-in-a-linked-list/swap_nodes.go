package leetcode

import "github.com/ympons/katas/essentials"

// ListNode represents a singly linked list node.
type ListNode = essentials.ListNode

func swapNodes(head *ListNode, k int) *ListNode {
	current := head
	for i := 1; i < k; i++ {
		current = current.Next
	}
	kNode, nkNode := current, head
	for current.Next != nil {
		nkNode = nkNode.Next
		current = current.Next
	}
	kNode.Val, nkNode.Val = nkNode.Val, kNode.Val
	return head
}
