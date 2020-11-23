package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ympons/katas/essentials"
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

var tests = []struct {
	l1, l2, expected []int
}{
	{
		l1:       []int{7, 2, 4, 3},
		l2:       []int{5, 6, 4},
		expected: []int{7, 8, 0, 7},
	},
	{
		l1:       []int{9},
		l2:       []int{1, 9},
		expected: []int{2, 8},
	},
}
