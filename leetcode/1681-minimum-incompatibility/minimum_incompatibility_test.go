package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumIncompatibility(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minimumIncompatibility(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		nums:     []int{1, 2, 1, 4},
		k:        2,
		expected: 4,
	},
	{
		nums:     []int{6, 3, 8, 1, 3, 1, 2, 2},
		k:        4,
		expected: 6,
	},
	{
		nums:     []int{5, 3, 3, 6, 3, 3},
		k:        3,
		expected: -1,
	},
	{
		nums:     []int{7, 3, 16, 15, 1, 13, 1, 2, 14, 5, 3, 10, 6, 2, 7, 15},
		k:        8,
		expected: 12,
	},
}
