package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertBST(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, convertBST(test.input))
	}
}

var tests = []struct {
	input    *TreeNode
	expected *TreeNode
}{
	{
		input: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		},
		expected: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 2},
		},
	},
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 4},
		},
		expected: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:  9,
				Left: &TreeNode{Val: 10},
			},
			Right: &TreeNode{Val: 4},
		},
	},
}
