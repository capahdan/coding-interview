package index_first_occur_string_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/28.index_first_occur_string"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Input
	expected int
}

type Input struct {
	haystack string
	needle   string
}

func TestLongestCommonPrefix(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    Input{haystack: "sadbutnotsad", needle: "sad"},
			expected: 0,
		},
		{
			name:     "test2",
			input:    Input{haystack: "leetcode", needle: "leeto"},
			expected: -1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, StrStr(tt.input.haystack, tt.input.needle))
		})
	}
}
