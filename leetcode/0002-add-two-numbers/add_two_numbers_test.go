package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/challenges/essentials"
)

func TestAddTwoNumbers(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := addTwoNumbers(
			essentials.IntsToList(test.l1),
			essentials.IntsToList(test.l2),
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

func TestAddTwoNumbersOld(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := addTwoNumbersOld(
			essentials.IntsToList(test.l1),
			essentials.IntsToList(test.l2),
		)
		assert.Equal(test.expected, essentials.ListToInts(output))
	}
}

var tests = []struct {
	l1       []int
	l2       []int
	expected []int
}{
	{
		l1:       []int{2, 4, 3},
		l2:       []int{5, 6, 4},
		expected: []int{7, 0, 8},
	},
	{
		l1:       []int{0},
		l2:       []int{0},
		expected: []int{0},
	},
	{
		l1:       []int{9, 9, 9, 9, 9, 9, 9},
		l2:       []int{9, 9, 9, 9},
		expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}
