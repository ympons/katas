package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := isBalanced(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected bool
}{
	{
		input:    nil,
		expected: true,
	},
	{
		input: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		expected: true,
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		},
		expected: false,
	},
}
