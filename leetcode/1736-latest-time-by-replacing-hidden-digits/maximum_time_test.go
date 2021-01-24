package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumTime(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, maximumTime(test.input))
	}
}

var tests = []struct {
	input    string
	expected string
}{
	{
		input:    "2?:?0",
		expected: "23:50",
	},
	{
		input:    "0?:3?",
		expected: "09:39",
	},
	{
		input:    "??:??",
		expected: "23:59",
	},
	{
		input:    "1?:22",
		expected: "19:22",
	},
}
