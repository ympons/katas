package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAncestorDiff(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxAncestorDiff(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected int
}{
	{
		input: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		expected: 3,
	},
	{
		input: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Right: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val: 14,
					Left: &TreeNode{
						Val: 13,
					},
				},
			},
		},
		expected: 7,
	},
}
