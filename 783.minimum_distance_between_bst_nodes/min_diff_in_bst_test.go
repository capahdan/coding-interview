package min_dif_in_bst_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/783.minimum_distance_between_bst_nodes"
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
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			output: 1,
		},
		{
			message: "test 2",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 48,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 49,
					},
				},
			},
			output: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, GetMinimumDifference(tt.input))
		})
	}

}
