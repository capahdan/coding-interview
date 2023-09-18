package length_of_last_word_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/58.length_of_last_word"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    string
	expected int
}

func TestLengthOfLastWord(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	name:     "test1",
		// 	input:    "Hello World",
		// 	expected: 5,
		// },
		{
			name:     "test2",
			input:    "    fly me   to   the moon    ",
			expected: 4,
		},
		// {
		// 	name:     "test3",
		// 	input:    "luffy is still joyboy",
		// 	expected: 6,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, LenghtOfLastWord(tt.input))
		})
	}
}
