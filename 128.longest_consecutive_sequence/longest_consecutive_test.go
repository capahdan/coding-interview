package longest_consecutive_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/128.longest_consecutive_sequence"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name   string
	input  []int
	output int
}

func TestLongestConsecutive(t *testing.T) {
	testTable := []TestTable{
		{
			name:   "test 1",
			input:  []int{100, 4, 200, 1, 3, 2},
			output: 4,
		},
		// {
		// 	name:   "test 2",
		// 	input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		// 	output: 9,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, LongestConsecutive(tt.input))
		})
	}
}
