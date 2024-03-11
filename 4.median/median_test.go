package median_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/4.median"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   input
	output  float64
}

type input struct {
	nums1 []int
	nums2 []int
}

func TestMedian(t *testing.T) {
	testTable := []testTable{
		{
			message: "test 1",
			input: input{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			output: 2.00000,
		},
		{
			message: "test 2",
			input: input{
				nums1: []int{1, 3},
				nums2: []int{2, 4},
			},
			output: 2.50000,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, Median(tt.input.nums1, tt.input.nums2))
		})
	}

}
