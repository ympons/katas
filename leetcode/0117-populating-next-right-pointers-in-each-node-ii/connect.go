package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	first := root
	dummy := &Node{}
	pre := dummy
	for first != nil {
		for first != nil {
			if first.Left != nil {
				pre.Next = first.Left
				pre = pre.Next
			}
			if first.Right != nil {
				pre.Next = first.Right
				pre = pre.Next
			}
			first = first.Next
		}
		first = dummy.Next
		dummy.Next = nil
		pre = dummy
	}
	return root
}
