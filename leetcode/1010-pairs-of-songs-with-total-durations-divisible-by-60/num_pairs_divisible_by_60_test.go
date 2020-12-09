package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := numPairsDivisibleBy60(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{30, 20, 150, 100, 40},
		expected: 3,
	},
	{
		input:    []int{60, 60, 60},
		expected: 3,
	},
}
