package number_of_one_bits_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/191.number_of_one_bits"
)

type TestTable struct {
	message string
	input   uint32
	output  int
}

func TestAddBinary(t *testing.T) {

	testTable := []TestTable{
		{
			message: "test 1",
			input:   11,
			output:  3,
		},
		// {
		// 	message: "test 2",
		// 	input:   128,
		// 	output:  1,
		// },
		// {
		// 	message: "test 3",
		// 	input:   4294967293,
		// 	output:  31,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, HammingWeight(tt.input))
		})
	}
}
