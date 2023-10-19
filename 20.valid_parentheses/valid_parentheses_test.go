package valid_parentheses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/20.valid_parentheses"
)

type TestTable struct {
	name     string
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    "()",
			expected: true,
		},
		{
			name:     "test2",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "test3",
			input:    "(]",
			expected: false,
		},
		{
			name:     "test4",
			input:    "]",
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsValid(tt.input))
		})
	}
}
