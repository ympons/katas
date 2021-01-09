package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalMoney(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, totalMoney(test.input))
	}
}

func TestTotalMoneyDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, totalMoneyDP(test.input))
	}
}

var tests = []struct {
	input, expected int
}{
	{
		input:    1,
		expected: 1,
	},
	{
		input:    10,
		expected: 37,
	},
	{
		input:    20,
		expected: 96,
	},
	{
		input:    1000,
		expected: 74926,
	},
}
