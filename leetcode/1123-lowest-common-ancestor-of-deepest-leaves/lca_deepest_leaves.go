package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := helper(root)
	return node
}

func helper(root *TreeNode) (node *TreeNode, height int) {
	if root == nil {
		return nil, 0
	}

	nodeL, heightL := helper(root.Left)
	nodeR, heightR := helper(root.Right)

	if heightL > heightR {
		return nodeL, heightL + 1
	}

	if heightL < heightR {
		return nodeR, heightR + 1
	}

	return root, heightL + 1
}
