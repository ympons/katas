package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	answer := &TreeNode{}
	current := answer
	helper(&current, node)
	return answer.Right
}

func helper(root **TreeNode, node *TreeNode) {
	if node.Left != nil {
		helper(root, node.Left)
	}
	if node != nil {
		(*root).Right = &TreeNode{Val: node.Val}
		*root = (*root).Right
	}
	if node.Right != nil {
		helper(root, node.Right)
	}
}
