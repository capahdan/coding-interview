package topk_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/347.topk"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   input
	output  []int
}

type input struct {
	nums []int
	k    int
}

func TestMaxArea(t *testing.T) {

	testTable := []testTable{
		// {
		// 	message: "test 1",
		// 	input: input{

		// 		nums: []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 4},
		// 		k:    2,
		// 	},
		// 	output: []int{1, 3},
		// },
		{
			message: "test 2",
			input: input{

				nums: []int{1},
				k:    1,
			},
			output: []int{1},
		},
		// {
		// 	message: "test 3",
		// 	input: input{

		// 		nums: []int{1, 1, 1, 2, 2, 3},
		// 		k:    2,
		// 	},
		// 	output: []int{1, 2},
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, TopKFrequent(tt.input.nums, tt.input.k))
		})
	}

}
