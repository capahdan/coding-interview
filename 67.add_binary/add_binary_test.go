package add_binary_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/67.add_binary"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  string
}

type input struct {
	a string
	b string
}

func TestAddBinary(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: input{
				a: "11",
				b: "1",
			},
			output: "100",
		},
		{
			message: "test 2",
			input: input{
				a: "1010",
				b: "1011",
			},
			output: "10101",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, AddBinary(tt.input.a, tt.input.b))
		})
	}
}
