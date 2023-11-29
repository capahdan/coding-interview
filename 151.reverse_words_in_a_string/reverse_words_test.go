package reverse_word_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/151.reverse_words_in_a_string"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   string
	output  string
}

func TestReverseWord(t *testing.T) {
	testTable := []TestTable{
		// {
		// 	message: "test1",
		// 	input:   "the sky is blue",
		// 	output:  "blue is sky the",
		// },
		{
			message: "test2",
			input:   "  hello     world  ",
			output:  "world hello",
		},
		// {
		// 	message: "test3",
		// 	input:   "a good   example",
		// 	output:  "example good a",
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, ReverseWord(tt.input))
		})
	}
}
