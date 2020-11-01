package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := letterCombinations(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected []string
}{
	{
		input:    "23",
		expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		input:    "",
		expected: []string{},
	},
	{
		input:    "2",
		expected: []string{"a", "b", "c"},
	},
}
