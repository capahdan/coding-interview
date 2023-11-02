package symmetric_tree_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/101.symmetric_tree"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   *TreeNode
	output  bool
}

func TestIsSameTree(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	message: "test 1",
		// 	input: &TreeNode{
		// 		Val: 1,
		// 		Left: &TreeNode{
		// 			Val: 2,
		// 			Left: &TreeNode{
		// 				Val: 3,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 4,
		// 			},
		// 		},
		// 		Right: &TreeNode{
		// 			Val: 2,
		// 			Left: &TreeNode{
		// 				Val: 4,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 			},
		// 		},
		// 	},
		// 	output: true,
		// },
		{
			message: "test 2",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			output: false,
		},
		// {
		// 	message: "test 3",
		// 	input: &TreeNode{
		// 		Val: 1,
		// 		Left: &TreeNode{
		// 			Val: 0,
		// 		},
		// 	},
		// 	output: false,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, IsSymmetric(tt.input))
		})
	}

}
