package average_of_level_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/637.average_of_levels_in_binary_tree"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   *TreeNode
	output  []float64
}

func TestMinDiffInBst(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
			output: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			message: "test 2",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			output: []float64{3.00000, 14.50000, 11.00000},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, AverageOfLevels(tt.input))
		})
	}

}
