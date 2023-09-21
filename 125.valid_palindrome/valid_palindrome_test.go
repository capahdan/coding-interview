package valid_palindrome_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/125.valid_palindrome"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    string
	expected bool
}

func TestValidPalindrome(t *testing.T) {
	testTable := []TestTable{

		{
			name:     "test1",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "test2",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "test3",
			input:    " ",
			expected: true,
		},
		{
			name:     "test4",
			input:    "a k h",
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ValidPalindrome(tt.input))
		})
	}
}
