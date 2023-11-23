package count_nodes_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/222.count_complete_tree_nodes"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   *TreeNode
	output  int
}

func TestMinDiffInBst(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			output: 6,
		},
		{
			message: "test 2",
			input:   nil,
			output:  0,
		},
		{
			message: "test 3",
			input: &TreeNode{
				Val: 1,
			},
			output: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, CountNodes(tt.input))
		})
	}

}
