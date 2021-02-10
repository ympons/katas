package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimBST(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, trimBST(test.root, test.low, test.high))
	}
}

var tests = []struct {
	root      *TreeNode
	low, high int
	expected  *TreeNode
}{
	{
		root: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		},
		low:  1,
		high: 2,
		expected: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
	},
	{
		root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: 4},
		},
		low:  1,
		high: 3,
		expected: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
	},
}
