package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := romanToInt(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "III",
		expected: 3,
	},
	{
		input:    "IV",
		expected: 4,
	},
	{
		input:    "IX",
		expected: 9,
	},
	{
		input:    "LVIII",
		expected: 58,
	},
	{
		input:    "MCMXCIV",
		expected: 1994,
	},
}
