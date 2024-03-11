package multiply_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/43.multiply_string"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   input
	output  string
}

type input struct {
	num1 string
	num2 string
}

func TestMultiply(t *testing.T) {

	testtable := []testTable{
		// {
		// 	message: "test 10 times 10 ",
		// 	input: input{
		// 		num1: "10",
		// 		num2: "11",
		// 	},
		// 	output: "110",
		// },
		// {
		// 	message: "test 123 times 456 ",
		// 	input: input{
		// 		num1: "123",
		// 		num2: "456",
		// 	},
		// 	output: "56088",
		// },
		{
			message: "test 2 times 3 ",
			input: input{
				num1: "2",
				num2: "3",
			},
			output: "6",
		},

		// {
		// 	message: "test 498828660196 times 840477629533 ",
		// 	input: input{
		// 		num1: "498828660196",
		// 		num2: "840477629533",
		// 	},
		// 	output: "419254329864656431168468",
		// },
	}

	for _, tt := range testtable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, MultiplyString(tt.input.num1, tt.input.num2))
		})
	}

}
