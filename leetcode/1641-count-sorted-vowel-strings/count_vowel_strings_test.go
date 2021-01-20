package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowelStrings(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, countVowelStrings(test.input))
	}
}

func TestCountVowelStringsDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, countVowelStringsDP(test.input))
	}
}

var tests = []struct {
	input    int
	expected int
}{
	{
		input:    1,
		expected: 5,
	},
	{
		input:    2,
		expected: 15,
	},
	{
		input:    3,
		expected: 35,
	},
	{
		input:    33,
		expected: 66045,
	},
}
