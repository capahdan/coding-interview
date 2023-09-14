package string_to_integer_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/string_to_integer"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    string
	expected int
}

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	name:     "test 1",
		// 	input:    "42",
		// 	expected: 42,
		// },
		// {
		// 	name:     "test 2",
		// 	input:    "41  ini test",
		// 	expected: 41,
		// },
		{
			name:     "test 3",
			input:    "i 987",
			expected: 987,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MyAtoi(tt.input))
		})
	}
}
