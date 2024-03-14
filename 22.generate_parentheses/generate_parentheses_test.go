package generateparentheses_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/22.generate_parentheses"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   int
	output  []string
}

func TestGenerateParentheses(t *testing.T) {
	testTable := []testTable{
		{
			message: "test n = 3",
			input:   3,
			output:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			message: "test n = 1",
			input:   1,
			output:  []string{"()"},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, GenerateParenthesis(tt.input))
		})
	}
}
