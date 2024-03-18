package reverselinkedlist_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/203.reverselinkedlist"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	Name   string
	Input  *ListNode
	Output *ListNode
}

func TestReverseList(t *testing.T) {

	testTable := []TestTable{
		{
			Name: "Test 1",
			Input: &ListNode{
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
			Output: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Output, ReverseList(tt.Input))
		})
	}
}
