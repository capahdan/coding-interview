package summary_range_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/228.summary_ranges"
)

type TestTable struct {
	name     string
	input    []int
	expected []string
}

func TestSummaryRanges(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    []int{0, 1, 2, 4, 5, 7},
			expected: []string{"0->2", "4->5", "7"},
		},
		{
			name:     "test2",
			input:    []int{0, 2, 3, 4, 6, 8, 9},
			expected: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, SummaryRanges(tt.input))
		})
	}
}
