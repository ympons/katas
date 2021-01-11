package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTimeRequiredDP(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minimumTimeRequiredDP(test.jobs, test.k)
		assert.Equal(test.expected, output)
	}
}

func TestMinimumTimeRequired(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minimumTimeRequired(test.jobs, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	jobs     []int
	k        int
	expected int
}{
	{
		jobs:     []int{3, 2, 3},
		k:        3,
		expected: 3,
	},
	{
		jobs:     []int{1, 2, 4, 7, 8},
		k:        2,
		expected: 11,
	},
}
