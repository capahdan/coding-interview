package merge_sorted_array_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/88.merge_sorted_array"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected []int
}

type Number struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "test2",
			input: Number{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			expected: []int{1},
		},
		{
			name: "test2",
			input: Number{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			expected: []int{1},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Merge(tt.input.nums1, tt.input.m, tt.input.nums2, tt.input.n))
		})
	}
}
