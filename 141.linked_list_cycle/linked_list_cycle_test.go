package valid_parentheses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/141.linked_list_cycle"
)

type TestTable struct {
	name     string
	input    ListNode
	expected bool
}

func TestIsValid(t *testing.T) {

	// Membuat linked list dengan cycle
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node2 // Membuat siklus dari node5 kembali ke node2

	a2 := &ListNode{Val: 2, Next: nil}
	a1 := &ListNode{Val: 1, Next: nil}

	a1.Next = a2
	a2.Next = a1

	testTable := []TestTable{
		// {
		// 	name:     "test1",
		// 	input:    *node1,
		// 	expected: true,
		// },
		// {
		// 	name:     "test2",
		// 	input:    *a1,
		// 	expected: true,
		// },
		// {
		// 	name: "test3",
		// 	input: ListNode{
		// 		Val:  1,
		// 		Next: nil,
		// 	},
		// 	expected: false,
		// },
		{
			name:     "test4",
			input:    ListNode{},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, HasCycle(&tt.input))
		})
	}
}
