package valid_sudoku_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/36.valid_sudoku"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   [][]byte
	output  bool
}

func TestValidSudoku(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			output: true,
		},
		{
			message: "test 2",
			input: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			output: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, IsValidSudoku(tt.input))
		})
	}
}
