package reorder_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/143.reorder_list"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   *ListNode
	output  *ListNode
}

func TestReorderList(t *testing.T) {

	testTable := []TestTable{
		{
			message: "test 1",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, ReorderList(tt.input))
		})
	}
}
