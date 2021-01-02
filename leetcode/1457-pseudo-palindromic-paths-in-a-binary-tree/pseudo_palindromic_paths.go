package leetcode

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	counter, answer := make([]int, 10), 0
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		counter[root.Val]++
		if root.Left == nil && root.Right == nil {
			if palindromic(counter) {
				answer++
			}
			return
		}
		if root.Left != nil {
			preorder(root.Left)
			counter[root.Left.Val]--
		}
		if root.Right != nil {
			preorder(root.Right)
			counter[root.Right.Val]--
		}
	}
	preorder(root)
	return answer
}

func palindromic(counter []int) bool {
	odd := 0
	for _, v := range counter {
		if v&1 == 1 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}
