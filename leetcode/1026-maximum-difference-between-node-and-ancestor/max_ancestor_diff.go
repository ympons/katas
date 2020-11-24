package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, root.Val, root.Val)
}

func helper(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	left := helper(root.Left, max, min)
	right := helper(root.Right, max, min)
	if left > right {
		return left
	}
	return right
}
