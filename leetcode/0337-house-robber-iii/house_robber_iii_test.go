package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := rob(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected int
}{
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		expected: 7,
	},
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		expected: 9,
	},
}
