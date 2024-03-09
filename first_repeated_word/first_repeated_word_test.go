package first_repeated_word_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/first_repeated_word"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message  string
	input    string
	expected string
}

func Test(t *testing.T) {
	testTable := []testTable{
		{
			message:  "test1",
			input:    "We work hard because hard work pays",
			expected: "hard",
		},
		{
			message:  "test2",
			input:    "We 	work, hard:because-hard.work pays",
			expected: "hard",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.expected, FirstRepeatedWord(tt.input))
		})
	}
}
