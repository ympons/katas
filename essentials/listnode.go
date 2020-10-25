package essentials

// ListNode represents a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// IntsToList converts a slice of ints to a linked list
func IntsToList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	head := new(ListNode)
	current := head
	for _, v := range slice {
		current.Next = &ListNode{v, nil}
		current = current.Next
	}

	return head.Next
}

// ListToInts converts a linked list to a slice of ints
func ListToInts(head *ListNode) []int {
	slice := []int{}
	for node := head; node != nil; node = node.Next {
		slice = append(slice, node.Val)
	}
	return slice
}
