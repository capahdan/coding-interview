package longest_substring_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/longest_palindromic_substring"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	test := []struct {
		names    string
		input    string
		expected string
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
			input:    "babad",
			expected: "bab",
		},
		{
			names:    "test 3",
			input:    "abbd",
			expected: "bb",
		},
		{
			names:    "test 3",
			input:    "aa",
			expected: "aa",
		},
		{
			names:    "test 3",
			input:    "abb",
			expected: "bb",
		},
	}

	for _, tt := range test {
		t.Run(tt.names, func(t *testing.T) {
			assert.Equal(t, tt.expected, LongestPalindrom(tt.input))
		})
	}
}
