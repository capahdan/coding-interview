package longest_substring_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/longest_substring_without_repeating_character"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	test := []struct {
		names    string
		input    string
		expected int
	}{
		// {
		// 	names:    "test 1",
		// 	input:    "abcabcbb",
		// 	expected: 3,
		// },
		// {
		// 	names:    "test 2",
		// 	input:    "bbbbb",
		// 	expected: 1,
		// },
		{
			names:    "test 3",
			input:    "pwwkew",
			expected: 3,
		},
	}

	for _, tt := range test {
		t.Run(tt.names, func(t *testing.T) {
			assert.Equal(t, tt.expected, LengthOfLongestSubstring(tt.input))
		})
	}
}
