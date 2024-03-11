package contain_duplicate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/217.contains_duplicates"
)

type TestTable struct {
	name     string
	input    []int
	expected bool
}

func TestContainDuplicate(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "test2",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ContainsDuplicate(tt.input))
		})
	}
}
