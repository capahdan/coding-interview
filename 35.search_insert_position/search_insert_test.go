package search_insert_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/35.search_insert_position"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  int
}

type input struct {
	nums   []int
	target int
}

func TestSearchInsert(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			output: 2,
		},
		{
			message: "test 2",
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			output: 1,
		},
		{
			message: "test 1",
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			output: 4,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, SearchInsert(tt.input.nums, tt.input.target))
		})
	}
}
