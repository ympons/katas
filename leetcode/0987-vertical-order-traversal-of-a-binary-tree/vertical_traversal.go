package leetcode

import "sort"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type tuple struct {
	row, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	M := map[int][]tuple{}
	colMax, colMin := -1001, 1001

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		if node == nil {
			return
		}
		M[c] = append(M[c], tuple{r, node.Val})
		if c > colMax {
			colMax = c
		}
		if c < colMin {
			colMin = c
		}
		dfs(node.Left, r+1, c-1)
		dfs(node.Right, r+1, c+1)
	}
	dfs(root, 0, 0)

	answer := [][]int{}
	for col := colMin; col <= colMax; col++ {
		answer = append(answer, []int{})
		sort.Slice(M[col], func(i, j int) bool {
			if M[col][i].row < M[col][j].row {
				return true
			}
			if M[col][i].row > M[col][j].row {
				return false
			}
			return M[col][i].val < M[col][j].val
		})
		for _, t := range M[col] {
			answer[len(answer)-1] = append(answer[len(answer)-1], t.val)
		}
	}
	return answer
}
