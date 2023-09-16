package majority_element_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/169.majority_element"
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

func TestMajorityElement(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums: []int{3, 2, 3},
			},
			expected: 3,
		},
		{
			name: "test2",
			input: Number{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			expected: 2,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MajorityElement(tt.input.nums))
		})
	}
}
