package three_sum_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/15.three_sum"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    []int
	expected [][]int
}

func TestTreeSum(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "test2",
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "test3",
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ThreeSum(tt.input))
		})
	}
}
