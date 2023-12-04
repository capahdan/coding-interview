package gas_station_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/134.gas_station"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name   string
	input  input
	output int
}

type input struct {
	gas  []int
	cost []int
}

func TestGasStation(t *testing.T) {
	testTable := []testTable{
		{
			name: "Test 1",
			input: input{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			output: 3,
		},
		// {
		// 	name: "Test 2",
		// 	input: input{
		// 		gas:  []int{2, 3, 4},
		// 		cost: []int{3, 4, 3},
		// 	},
		// 	output: -1,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, CanCompleteCircuit(tt.input.gas, tt.input.cost))
		})
	}
}
