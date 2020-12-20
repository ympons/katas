package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReformatNumber(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := reformatNumber(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input, expected string
}{
	{
		input:    "1-23-45 6",
		expected: "123-456",
	},
	{
		input:    "123 4-567",
		expected: "123-45-67",
	},
	{
		input:    "123 4-5678",
		expected: "123-456-78",
	},
	{
		input:    "--17-5 229 35-39475 ",
		expected: "175-229-353-94-75",
	},
	{
		input:    "123",
		expected: "123",
	},
}
