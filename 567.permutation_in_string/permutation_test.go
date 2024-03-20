package permutation_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/567.permutation_in_string"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   input
	output  bool
}

type input struct {
	s1 string
	s2 string
}

func TestCheckInclussion(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test1",
			input: input{
				s1: "ab",
				s2: "eidbaooo",
			},
			output: true,
		},
		{
			message: "test1",
			input: input{
				s1: "ab",
				s2: "eidboaoo",
			},
			output: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, CheckInclusion(tt.input.s1, tt.input.s2))
		})
	}
}
