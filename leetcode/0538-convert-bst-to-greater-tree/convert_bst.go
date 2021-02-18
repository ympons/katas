package leetcode

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var (
		sum     int
		inorder func(*TreeNode)
	)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		sum += node.Val
		node.Val = sum
		inorder(node.Left)
	}
	inorder(root)
	return root
}
