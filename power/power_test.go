package power_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/power"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	input   input
	message string
	output  float32
}

type input struct {
	a float32
	b int
}

func TestPower(t *testing.T) {
	TestTable := []TestTable{
		{
			input:   input{a: 2, b: 2},
			output:  4,
			message: "test1",
		},
	}

	for _, tt := range TestTable {

		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, Power(tt.input.a, tt.input.b))
		})
	}

}
