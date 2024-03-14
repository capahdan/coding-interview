package two_sum_2_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/167.two_sum_2"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected []int
}

type Number struct {
	nums   []int
	target int
}

func TestTwoSum2(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{1, 2},
		},
		{
			name: "test2",
			input: Number{
				nums:   []int{2, 3, 4},
				target: 6,
			},
			expected: []int{1, 3},
		},
		{
			name: "test3",
			input: Number{
				nums:   []int{-1, 0},
				target: -1,
			},
			expected: []int{1, 2},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TwoSum2(tt.input.nums, tt.input.target))
		})
	}
}
