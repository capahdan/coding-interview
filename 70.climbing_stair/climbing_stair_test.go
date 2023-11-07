package climbing_stair_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/70.climbing_stair"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    int
	expected int
}

func TestClimbingStairs(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	name:     "test 1",
		// 	input:    2,
		// 	expected: 2,
		// },
		// {
		// 	name:     "test 2",
		// 	input:    3,
		// 	expected: 3,
		// },
		{
			name:     "test 4",
			input:    4,
			expected: 5,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ClimbingStairs(tt.input))
		})
	}
}
