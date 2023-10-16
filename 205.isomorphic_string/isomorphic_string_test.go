package isomorphic_string_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/205.isomorphic_string"
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

func TestIsomorphicString(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	name: "test1",
		// 	input: input{
		// 		s: "egg",
		// 		t: "add",
		// 	},
		// 	expected: true,
		// },
		{
			name: "test2",
			input: input{
				s: "foo",
				t: "bar",
			},
			expected: false,
		},
		// {
		// 	name: "test3",
		// 	input: input{
		// 		s: "paper",
		// 		t: "title",
		// 	},
		// 	expected: true,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsomorphicString(tt.input.s, tt.input.t), tt.expected)
		})
	}
}
