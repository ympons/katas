package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumBST(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := rangeSumBST(test.root, test.low, test.high)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	root      *TreeNode
	low, high int
	expected  int
}{
	{
		root: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 15,
				Right: &TreeNode{
					Val: 18,
				},
			},
		},
		low:      7,
		high:     15,
		expected: 32,
	},
}
