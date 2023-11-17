package path_sum_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/112.path_sum"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  bool
}

type input struct {
	root      *TreeNode
	targetSum int
}

func TestTreeNode(t *testing.T) {
	testTable := []TestTable{
		{
			message: "Test 1",
			input: input{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				targetSum: 22,
			},
			output: true,
		},
		// {
		// 	message: "Test 2",
		// 	input: input{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 			Left: &TreeNode{
		// 				Val: 2,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 			},
		// 		},
		// 		targetSum: 5,
		// 	},
		// 	output: false,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, HasPathSum(tt.input.root, tt.input.targetSum))
		})
	}
}
