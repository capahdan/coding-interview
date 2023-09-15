package remove_element_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/27.remove_element"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected int
}

type Number struct {
	nums []int
	val  int
}

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			expected: 2,
		},
		{
			name: "test2",
			input: Number{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			expected: 5,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, RemoveElement(tt.input.nums, tt.input.val))
		})
	}
}
