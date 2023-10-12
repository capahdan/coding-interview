package valid_anagram_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/242.valid_anagram"
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

func TestLongestCommonPrefix(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    input{s: "anagram", t: "nagaram"},
			expected: true,
		},
		{
			name:     "test2",
			input:    input{s: "rat", t: "car"},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ValidAnagram(tt.input.s, tt.input.t))
		})
	}
}
