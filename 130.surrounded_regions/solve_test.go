package solve_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/130.surrounded_regions"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name   string
	input  [][]byte
	output [][]byte
}

func TestSolve(t *testing.T) {

	testTable := []testTable{
		{
			name: "test 1",
			input: [][]byte{
				{
					'X', 'X', 'X', 'X',
				},
				{
					'X', 'O', 'O', 'X',
				},
				{
					'X', 'X', 'O', 'X',
				},
				{
					'X', 'O', 'X', 'X',
				},
			},

			output: [][]byte{
				{
					'X', 'X', 'X', 'X',
				},
				{
					'X', 'X', 'X', 'X',
				},
				{
					'X', 'X', 'X', 'X',
				},
				{
					'X', 'O', 'X', 'X',
				},
			},
		},
		{
			name: "test 1",
			input: [][]byte{
				{
					'X',
				},
			},

			output: [][]byte{
				{
					'X',
				},
			},
		},
		{
			name: "test 1",
			input: [][]byte{
				{
					'X', 'O', 'X',
				},
				{
					'X', 'O', 'X',
				},
				{
					'X', 'O', 'X',
				},
			},

			output: [][]byte{
				{
					'X', 'O', 'X',
				},
				{
					'X', 'O', 'X',
				},
				{
					'X', 'O', 'X',
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, Solve(tt.input))
		})
	}
}
