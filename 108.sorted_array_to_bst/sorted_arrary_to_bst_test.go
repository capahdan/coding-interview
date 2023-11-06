package sorted_arr_to_bst_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/108.sorted_array_to_bst"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name   string
	input  []int
	output *TreeNode
}

func TestSortedArrayToBST(t *testing.T) {
	testTable := []testTable{
		// {
		// 	name:  "test 1",
		// 	input: []int{-10, -3, 0, 5, 9},
		// 	output: &TreeNode{
		// 		Val: 0,
		// 		Left: &TreeNode{
		// 			Val: -3,
		// 			Left: &TreeNode{
		// 				Val: -10,
		// 			},
		// 		},
		// 		Right: &TreeNode{
		// 			Val: 9,
		// 			Left: &TreeNode{
		// 				Val: 5,
		// 			},
		// 		},
		// 	},
		// },
		{
			name:  "test 2",
			input: []int{1, 3},
			output: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, SortedArrayToBST(tt.input))
		})
	}

}
