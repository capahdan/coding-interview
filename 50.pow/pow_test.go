package pow_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/50.pow"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   input
	output  float64
}

type input struct {
	x float64
	n int
}

func TestInterview(t *testing.T) {

	testtable := []testTable{
		{
			message: "test 2 ^ 10 ",
			input: input{
				x: 2.00000,
				n: 10,
			},
			output: 1024.00000,
		},
		{
			message: "test 2.1000  ^ 3",
			input: input{
				x: 2.10000,
				n: 3,
			},
			output: 9.26100,
		},
		{
			message: "test 2.00000 ^ -2",
			input: input{
				x: 2.00000,
				n: -2,
			},
			output: 0.25000,
		},
	}

	for _, tt := range testtable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, Power(tt.input.x, tt.input.n))
		})
	}

}
