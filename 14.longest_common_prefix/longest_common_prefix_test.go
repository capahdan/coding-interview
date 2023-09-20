package longest_common_prefix_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/14.longest_common_prefix"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "test2",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, LongestCommonPrefix(tt.input))
		})
	}
}
