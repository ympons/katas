package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressiveWords(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := expressiveWords(test.S, test.words)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	S        string
	words    []string
	expected int
}{
	{
		S:        "heeellooo",
		words:    []string{"hello", "hi", "helo"},
		expected: 1,
	},
	{
		S:        "abcd",
		words:    []string{"abc"},
		expected: 0,
	},
}
