package sum_number_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/129.sum_root_to_leaf_number"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   TreeNode
	output  int
}

func TestTreeNode(t *testing.T) {
	testTable := []TestTable{
		{
			message: "Test 1",
			input: TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			output: 25,
		},
		{
			message: "Test 2",
			input: TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			output: 1026,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, SumNumber(&tt.input))
		})
	}
}
