package max_depth_binary_tree_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/104.maximum_depth_binary_tree"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   TreeNode
	output  int
}

func TestTreeNode(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	message: "Test 1",
		// 	input: TreeNode{
		// 		Val: 3,
		// 		Left: &TreeNode{
		// 			Val: 9,
		// 		},
		// 		Right: &TreeNode{
		// 			Val: 20,
		// 			Left: &TreeNode{
		// 				Val: 15,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 7,
		// 			},
		// 		},
		// 	},
		// 	output: 3,
		// },
		{
			message: "Test 2",
			input: TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			output: 2,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, MaxDepth(&tt.input))
		})
	}
}
