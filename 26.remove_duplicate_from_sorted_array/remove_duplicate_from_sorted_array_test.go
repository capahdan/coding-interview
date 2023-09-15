package remove_duplicate_from_sorted_array_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/26.remove_duplicate_from_sorted_array"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected int
}

type Number struct {
	nums []int
}

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums: []int{1, 1, 2},
			},
			expected: 2,
		},
		{
			name: "test2",
			input: Number{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			expected: 5,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, RemoveDuplicates(tt.input.nums))
		})
	}
}
