package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPseudoPalindromicPaths(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := pseudoPalindromicPaths(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected int
}{
	{
		input: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
			Right: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
			},
		},
		expected: 2,
	},
	{
		input: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: 1},
		},
		expected: 1,
	},
	{
		input:    &TreeNode{Val: 8},
		expected: 1,
	},
}
