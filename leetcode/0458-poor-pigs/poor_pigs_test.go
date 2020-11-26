package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoorPigs(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := poorPigs(test.buckets, test.minutesToDie, test.minutesToTest)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	buckets, minutesToDie, minutesToTest int
	expected                             int
}{
	{
		buckets:       1000,
		minutesToDie:  15,
		minutesToTest: 60,
		expected:      5,
	},
	{
		buckets:       4,
		minutesToDie:  15,
		minutesToTest: 15,
		expected:      2,
	},
	{
		buckets:       4,
		minutesToDie:  15,
		minutesToTest: 30,
		expected:      2,
	},
}
