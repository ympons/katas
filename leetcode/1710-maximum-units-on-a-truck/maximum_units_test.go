package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumUnits(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maximumUnits(test.boxTypes, test.truckSize)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	boxTypes  [][]int
	truckSize int
	expected  int
}{
	{
		boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
		truckSize: 4,
		expected:  8,
	},
	{
		boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
		truckSize: 10,
		expected:  91,
	},
}
