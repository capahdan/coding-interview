package integer_to_roman_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/12.integer_to_roman"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   int
	output  string
}

func TestIntegerToRoman(t *testing.T) {
	testTable := []testTable{
		{
			message: "test1",
			input:   3,
			output:  "III",
		},
		{
			message: "test2",
			input:   58,
			output:  "LVIII",
		},
		{
			message: "test3",
			input:   1994,
			output:  "MCMXCIV",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, IntToRoman(tt.input))
		})
	}
}
