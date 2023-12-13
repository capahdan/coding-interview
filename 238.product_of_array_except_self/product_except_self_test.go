package product_except_self_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/238.product_of_array_except_self"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name   string
	input  []int
	output []int
}

func TestProductExceptSelf(t *testing.T) {
	testTable := []testTable{
		{
			name:   "test 1",
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			name:   "test 2",
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, ProductExceptSelf(tt.input))
		})
	}
}
