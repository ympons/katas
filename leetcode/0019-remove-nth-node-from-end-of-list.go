/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

/*
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
Given n will always be valid.
Try to do this in one pass.
*/