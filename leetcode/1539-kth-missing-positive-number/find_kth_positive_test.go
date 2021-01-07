package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthPositiveLog(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findKthPositiveLog(test.arr, test.k)
		assert.Equal(test.expected, output)
	}
}

func TestFindKthPositive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findKthPositive(test.arr, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	arr      []int
	k        int
	expected int
}{
	{
		arr:      []int{2, 3, 4, 7, 11},
		k:        5,
		expected: 9,
	},
	{
		arr:      []int{1, 2, 3, 4},
		k:        2,
		expected: 6,
	},
}
