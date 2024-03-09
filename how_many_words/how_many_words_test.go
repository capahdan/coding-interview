package how_many_words_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/how_many_words"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message  string
	input    input
	expected []int64
}

type input struct {
	wordSet   []string
	sentences []string
}

func Test(t *testing.T) {
	testTable := []testTable{
		{
			message: "test1",
			input: input{
				wordSet:   []string{"the", "bats", "tabs", "in", "cat", "act"},
				sentences: []string{"cat the bats", "in the act", "act tabs in"},
			},
			expected: []int64{4, 2, 4},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.expected, CountSentences(tt.input.wordSet, tt.input.sentences))
		})
	}
}
