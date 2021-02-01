package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticalTraversal(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, verticalTraversal(test.input))
	}
}

var tests = []struct {
	input    *TreeNode
	expected [][]int
}{
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
		},
		expected: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
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
		expected: [][]int{{9}, {3, 15}, {20}, {7}},
	},
}
