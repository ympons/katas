package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRescueBoats(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, numRescueBoats(test.people, test.limit))
	}
}

var tests = []struct {
	people   []int
	limit    int
	expected int
}{
	{
		people:   []int{1, 2},
		limit:    3,
		expected: 1,
	},
	{
		people:   []int{3, 2, 2, 1},
		limit:    3,
		expected: 3,
	},
	{
		people:   []int{3, 5, 3, 4},
		limit:    5,
		expected: 4,
	},
}
