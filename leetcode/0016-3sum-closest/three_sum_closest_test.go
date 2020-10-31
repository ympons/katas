package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := threeSumClosest(test.nums, test.target)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	target   int
	expected int
}{
	{
		nums:     []int{-1, 2, 1, -4},
		target:   1,
		expected: 2,
	},
}
