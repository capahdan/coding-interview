package word_pattern_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/290.word_pattern"
)

type TestTable struct {
	name     string
	input    input
	expected bool
}
type input struct {
	pattern string
	s       string
}

func TestWordPattern(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: input{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			expected: true,
		},
		{
			name: "test2",
			input: input{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			expected: false,
		},
		// {
		// 	name: "test1",
		// 	input: input{
		// 		pattern: "aaaa",
		// 		s:       "dog cat cat dog",
		// 	},
		// 	expected: false,
		// },
		{
			name: "test4",
			input: input{
				pattern: "abba",
				s:       "dog dog dog dog",
			},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, WordPattern(tt.input.pattern, tt.input.s), tt.expected)
		})
	}
}
