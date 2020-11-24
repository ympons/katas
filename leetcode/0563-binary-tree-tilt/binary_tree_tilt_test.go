package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTilt(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findTilt(test.input)
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
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		expected: 1,
	},
	{
		input: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		expected: 15,
	},
}
