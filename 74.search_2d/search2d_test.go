package search2d_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/74.search_2d"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  bool
}

type input struct {
	arr    [][]int
	target int
}

func TestSeach2D(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: input{
				arr:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 3,
			},
			output: true,
		},
		{
			message: "test 2",
			input: input{
				arr:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 13,
			},
			output: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, SearchMatrix(tt.input.arr, tt.input.target))
		})
	}
}
