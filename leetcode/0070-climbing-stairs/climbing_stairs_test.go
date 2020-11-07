package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := climbStairs(test.input)
		assert.Equal(test.expected, output)
	}
}

func TestClimbStairsDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := climbStairsDP(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected int
}{
	{
		input:    2,
		expected: 2,
	},
	{
		input:    3,
		expected: 3,
	},
}
