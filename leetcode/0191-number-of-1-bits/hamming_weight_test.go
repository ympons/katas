package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, hammingWeight(test.input))
	}
}

var tests = []struct {
	input    uint32
	expected int
}{
	{
		input:    uint32(0b00000000000000000000000000001011),
		expected: 3,
	},
	{
		input:    uint32(0b00000000000000000000000010000000),
		expected: 1,
	},
	{
		input:    uint32(0b11111111111111111111111111111101),
		expected: 31,
	},
}
