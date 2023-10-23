package valid_parentheses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/21.merge_two_sorted_list"
)

type TestTable struct {
	Name     string
	Input    Input
	Expected *ListNode
}

type Input struct {
	l1 *ListNode
	l2 *ListNode
}

func TestIsValid(t *testing.T) {

	testTable := []TestTable{
		{
			Name: "Test 1",
			Input: Input{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			Expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expected, MergeTwoLists(tt.Input.l1, tt.Input.l2))
		})
	}
}
