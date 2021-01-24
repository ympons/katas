package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCharacters(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, minCharacters(test.a, test.b))
	}
}

var tests = []struct {
	a, b     string
	expected int
}{
	{
		a:        "aba",
		b:        "caa",
		expected: 2,
	},
	{
		a:        "dabadd",
		b:        "cda",
		expected: 3,
	},
	{
		a:        "dee",
		b:        "a",
		expected: 0,
	},
	{
		a:        "e",
		b:        "cce",
		expected: 1,
	},
	{
		a:        "a",
		b:        "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		expected: 2,
	},
}
