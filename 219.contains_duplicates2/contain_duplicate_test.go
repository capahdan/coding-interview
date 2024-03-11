package contain_duplicate_2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/219.contains_duplicates2"
)

type TestTable struct {
	name     string
	input    input
	expected bool
}
type input struct {
	nums []int
	k    int
}

func TestContainDuplicate(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: input{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			expected: true,
		},
		{
			name: "test2",
			input: input{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			expected: true,
		},
		{
			name: "test3",
			input: input{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ContainsNearbyDuplicate(tt.input.nums, tt.input.k), tt.expected)
		})
	}
}
