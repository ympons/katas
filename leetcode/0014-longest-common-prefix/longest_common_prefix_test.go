package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := longestCommonPrefix(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		input:    []string{"dog", "racecar", "car"},
		expected: "",
	},
}
