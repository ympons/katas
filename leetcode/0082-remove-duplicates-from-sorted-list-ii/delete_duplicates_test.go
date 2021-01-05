package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
)

func TestDeleteDuplicates(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := deleteDuplicates(essentials.IntsToList(test.input))
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

func TestDeleteDuplicatesClone(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := deleteDuplicatesClone(essentials.IntsToList(test.input))
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	input, expected []int
}{
	{
		input:    []int{1, 2, 3, 3, 4, 4, 5},
		expected: []int{1, 2, 5},
	},
	{
		input:    []int{1, 1, 1, 2, 3},
		expected: []int{2, 3},
	},
	{
		input:    []int{},
		expected: []int{},
	},
	{
		input:    []int{1},
		expected: []int{1},
	},
	{
		input:    []int{1, 1},
		expected: []int{},
	},
}
