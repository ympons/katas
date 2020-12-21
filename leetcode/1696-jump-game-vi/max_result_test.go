package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxResult(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxResult(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		nums:     []int{1, -1, -2, 4, -7, 3},
		k:        2,
		expected: 7,
	},
	{
		nums:     []int{10, -5, -2, 4, 0, 3},
		k:        3,
		expected: 17,
	},
	{
		nums:     []int{1, -5, -20, 4, -1, 3, -6, -3},
		k:        2,
		expected: 0,
	},
}
