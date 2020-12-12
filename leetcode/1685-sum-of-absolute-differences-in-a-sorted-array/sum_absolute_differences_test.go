package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumAbsoluteDifferences(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := getSumAbsoluteDifferences(test.nums)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	expected []int
}{
	{
		nums:     []int{2, 3, 5},
		expected: []int{4, 3, 5},
	},
	{
		nums:     []int{1, 4, 6, 8, 10},
		expected: []int{24, 15, 13, 15, 21},
	},
}
