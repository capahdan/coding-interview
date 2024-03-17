package binarysearch_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/704.binary_search"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  int
}

type input struct {
	arr    []int
	target int
}

func TestBinarySearch(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test1",
			input: input{
				arr:    []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			output: 4,
		},
		{
			message: "test2",
			input: input{
				arr:    []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			output: -1,
		},
		{
			message: "test3",
			input: input{
				arr:    []int{-1, 0, 3, 5, 9, 12},
				target: 0,
			},
			output: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, BinarySearch(tt.input.arr, tt.input.target))
		})
	}
}
