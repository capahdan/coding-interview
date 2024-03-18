package plus_one_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/66.plus_one"
)

type TestTable struct {
	name     string
	input    []int
	expected []int
}

func TestPlusOne(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test 1",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "test 2",
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "test 3",
			input:    []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "test 4",
			input:    []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			expected: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
		{
			name:     "test 5",
			input:    []int{1, 0, 9},
			expected: []int{1, 1, 0},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, PlusOne(tt.input))
		})
	}
}
