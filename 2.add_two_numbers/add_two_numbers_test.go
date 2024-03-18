package add_two_number_test

import (
	"reflect"
	"testing"

	. "github.com/capahdan/coding-interview/2.add_two_numbers"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	assert.Equal(t, expected, AddTwoNumbers(l1, l2))

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	test := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "test case 1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("AddTwoNumbers()= %v,want %v", got, tt.expected)
			}
		})
	}
}
