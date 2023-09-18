package rotate_array_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/189.rotate_array"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected []int
}

type Number struct {
	nums []int
	k    int
}

func TestRotateArray(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "test2",
			input: Number{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: []int{3, 99, -1, -100},
		},
		{
			name: "test3",
			input: Number{
				nums: []int{-1},
				k:    2,
			},
			expected: []int{-1},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, RotateArray(tt.input.nums, tt.input.k))
		})
	}
}
