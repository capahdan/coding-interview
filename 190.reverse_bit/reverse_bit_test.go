package reverse_bit_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/190.reverse_bit"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   uint32
	output  uint32
}

func TestAddBinary(t *testing.T) {

	testTable := []TestTable{
		{
			message: "test 1",
			input:   43261596,
			output:  964176192,
		},
		// {
		// 	message: "test 2",
		// 	input:   4294967293,
		// 	output:  3221225471,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, ReverseBit(tt.input))
		})
	}
}
