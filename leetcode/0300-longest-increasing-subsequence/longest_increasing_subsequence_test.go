package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := lengthOfLIS(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestLengthOfLISOld(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := lengthOfLISOld(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{10, 9, 2, 5, 3, 7, 101, 18},
		expected: 4,
	},
	{
		input:    []int{0, 1, 0, 3, 2, 3},
		expected: 4,
	},
	{
		input:    []int{7, 7, 7, 7, 7, 7, 7},
		expected: 1,
	},
}
