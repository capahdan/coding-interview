package missingnumber_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/268.missing_number"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   []int
	output  int
}

func TestMissingNumber(t *testing.T) {
	TestTable := []TestTable{
		{
			message: "test 1",
			input:   []int{3, 0, 1},
			output:  2,
		},
		{
			message: "test 2",
			input:   []int{0, 1},
			output:  2,
		},
		{
			message: "test 3",
			input:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			output:  8,
		},
	}

	for _, tt := range TestTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, MissingNumber(tt.input))
		})
	}
}
