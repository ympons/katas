package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) (left, right int) {
	if root == nil {
		return 0, 0
	}

	l1, r1 := helper(root.Left)
	l2, r2 := helper(root.Right)

	left = root.Val + r1 + r2

	if l1 > r1 {
		right = l1
	} else {
		right = r1
	}

	if l2 > r2 {
		right += l2
	} else {
		right += r2
	}

	return
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := helper(root)
	if left > right {
		return left
	}
	return right
}
