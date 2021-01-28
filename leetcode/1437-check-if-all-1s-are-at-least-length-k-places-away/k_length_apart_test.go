package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKLengthApart(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, kLengthApart(test.nums, test.k))
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected bool
}{
	{
		nums:     []int{1, 0, 0, 0, 1, 0, 0, 1},
		k:        2,
		expected: true,
	},
	{
		nums:     []int{1, 0, 0, 1, 0, 1},
		k:        2,
		expected: false,
	},
	{
		nums:     []int{1, 1, 1, 1, 1},
		k:        0,
		expected: true,
	},
	{
		nums:     []int{0, 1, 0, 1},
		k:        1,
		expected: true,
	},
}
