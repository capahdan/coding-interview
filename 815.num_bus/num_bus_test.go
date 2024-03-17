package numbus_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/capahdan/coding-interview/815.num_bus"
)

type TestTable struct {
	message string
	input   input
	output  int
}

type input struct {
	routes [][]int
	source int
	target int
}

func TestNumBus(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input: input{
				routes: [][]int{{1, 2, 7}, {3, 6, 7}},
				source: 1,
				target: 6,
			},
			output: 2,
		},
		// {
		// 	message: "test 2",
		// 	input: input{
		// 		routes: [][]int{{2, 7}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
		// 		source: 15,
		// 		target: 12,
		// 	},
		// 	output: -1,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, NumBusesToDestination(tt.input.routes, tt.input.source, tt.input.target))
		})
	}
}
