package valid_palindrome_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/392.is_subsequence"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    input
	expected bool
}

type input struct {
	s string
	t string
}

func TestIsSubsequence(t *testing.T) {
	testTable := []TestTable{

		{
			name: "test1",
			input: input{
				s: "abc",
				t: "ahbgdc",
			},
			expected: true,
		},
		{
			name: "test1",
			input: input{
				s: "axc",
				t: "ahbgdc",
			},
			expected: false,
		},
		{
			name: "test1",
			input: input{
				s: "b",
				t: "abc",
			},
			expected: true,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsSubsequence(tt.input.s, tt.input.t))
		})
	}
}
