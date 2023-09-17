package two_sum_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/1.two_sum"
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

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "test2",
			input: Number{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			name: "test3",
			input: Number{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TwoSum(tt.input.nums, tt.input.target))
		})
	}
}
