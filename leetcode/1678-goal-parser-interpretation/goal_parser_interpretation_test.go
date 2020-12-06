package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := interpret(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected string
}{
	{
		input:    "G()(al)",
		expected: "Goal",
	},
	{
		input:    "G()()()()(al)",
		expected: "Gooooal",
	},
	{
		input:    "(al)G(al)()()G",
		expected: "alGalooG",
	},
}
