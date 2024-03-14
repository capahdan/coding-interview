package reversepolis_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/150.reverse_polis"
	"github.com/stretchr/testify/assert"
)

// ["2","1","+","3","*"]

// *, /, +, -

// ["4","13","5","/","+"]

// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]

type TestTable struct {
	message string
	input   []string
	output  int
}

func TestEvalRPN(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input:   []string{"2", "1", "+", "3", "*"},
			output:  9,
		},
		{
			message: "test 2",
			input:   []string{"4", "13", "5", "/", "+"},
			output:  6,
		},
		{
			message: "test 3",
			input:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			output:  22,
		},
	}

	for _, tt := range testTable {

		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, EvalRPN(tt.input))
		})
	}

}
