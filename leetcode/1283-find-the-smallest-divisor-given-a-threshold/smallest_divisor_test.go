package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestDivisor(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := smallestDivisor(test.nums, test.threshold)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums      []int
	threshold int
	expected  int
}{
	{
		nums:      []int{1, 2, 5, 9},
		threshold: 6,
		expected:  5,
	},
	{
		nums:      []int{2, 3, 5, 7, 11},
		threshold: 11,
		expected:  3,
	},
	{
		nums:      []int{19},
		threshold: 5,
		expected:  4,
	},
}
