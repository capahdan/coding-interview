package same_tree_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/100.same_tree"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  bool
}

type input struct {
	left  TreeNode
	right TreeNode
}

func TestIsSameTree(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	message: "Test 1",
		// 	input: input{
		// 		left: TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 			},
		// 		},
		// 		right: TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 			},
		// 		},
		// 	},
		// 	output: true,
		// },
		{
			message: "Test 2",
			input: input{
				left: TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				right: TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			output: false,
		},
		// {
		// 	message: "Test 3",
		// 	input: input{
		// 		left: TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 1,
		// 			},
		// 		},
		// 		right: TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 1,
		// 			},
		// 		},
		// 	},
		// 	output: false,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, IsSameTree(&tt.input.left, &tt.input.right))
		})
	}
}
