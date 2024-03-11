package container_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/11.container"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   []int
	output  int
}

func TestMaxArea(t *testing.T) {

	testTable := []testTable{
		{
			message: "test 1",
			input:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output:  49,
		},
		{
			message: "test 2",
			input:   []int{1, 1},
			output:  1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, MaxArea(tt.input))
		})
	}

}
