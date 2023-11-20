package single_number_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/136.single_number"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name   string
	input  []int
	output int
}

func TestSingleNumber(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	name:   "test1",
		// 	input:  []int{2, 2, 1},
		// 	output: 1,
		// },
		{
			name:   "test2",
			input:  []int{4, 1, 2, 1, 2},
			output: 4,
		},
		{
			name:   "test3",
			input:  []int{1},
			output: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, SingleNumber(tt.input))
		})
	}
}
