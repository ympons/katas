package leetcode

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	answer := true
	helper(root, &answer)
	return answer
}

func helper(root *TreeNode, answer *bool) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, answer)
	right := helper(root.Right, answer)
	if abs(left-right) > 1 {
		*answer = false
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
