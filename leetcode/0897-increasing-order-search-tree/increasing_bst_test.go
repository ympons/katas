package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingBST(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := increasingBST(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    *TreeNode
	expected *TreeNode
}{
	{
		input: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 7},
		},
		expected: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	},
	{
		input: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
		},
		expected: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
								Right: &TreeNode{
									Val: 7,
									Right: &TreeNode{
										Val: 8,
										Right: &TreeNode{
											Val: 9,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
