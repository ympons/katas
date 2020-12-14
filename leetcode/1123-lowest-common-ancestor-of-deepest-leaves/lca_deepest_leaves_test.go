package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLcaDeepestLeaves(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := lcaDeepestLeaves(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected *TreeNode
}{
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		expected: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	},
}
