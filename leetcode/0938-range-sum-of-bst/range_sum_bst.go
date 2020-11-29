package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}
	if low < root.Val && root.Left != nil {
		sum += rangeSumBST(root.Left, low, high)
	}
	if high > root.Val && root.Left != nil {
		sum += rangeSumBST(root.Right, low, high)
	}
	return sum
}
