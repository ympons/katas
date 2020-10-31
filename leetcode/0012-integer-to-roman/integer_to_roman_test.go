package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := intToRoman(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    int
	expected string
}{
	{
		input:    3,
		expected: "III",
	},
	{
		input:    4,
		expected: "IV",
	},
	{
		input:    9,
		expected: "IX",
	},
	{
		input:    58,
		expected: "LVIII",
	},
	{
		input:    1994,
		expected: "MCMXCIV",
	},
}
