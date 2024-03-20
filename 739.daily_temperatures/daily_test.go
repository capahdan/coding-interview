package daily_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/739.daily_temperatures"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	message string
	input   []int
	output  []int
}

func TestDailyTemperatures(t *testing.T) {
	testTable := []TestTable{
		{
			message: "test 1",
			input:   []int{73, 74, 75, 71, 69, 72, 76, 73},
			output:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			message: "test 2",
			input:   []int{30, 40, 50, 50},
			output:  []int{1, 1, 1, 0},
		},
		{
			message: "test 3",
			input:   []int{30, 60, 90},
			output:  []int{1, 1, 0},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, DailyTemperatures(tt.input))
		})
	}
}
