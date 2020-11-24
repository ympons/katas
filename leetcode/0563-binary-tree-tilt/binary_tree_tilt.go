package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return abs(sum(root.Left)-sum(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}
