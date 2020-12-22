package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRangeII(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := smallestRangeII(test.A, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	A        []int
	k        int
	expected int
}{
	{
		A:        []int{1},
		k:        0,
		expected: 0,
	},
	{
		A:        []int{0, 10},
		k:        2,
		expected: 6,
	},
	{
		A:        []int{1, 3, 6},
		k:        3,
		expected: 3,
	},
}
