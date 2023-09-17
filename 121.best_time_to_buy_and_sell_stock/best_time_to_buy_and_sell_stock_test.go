package best_time_to_buy_and_sell_stock_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/121.best_time_to_buy_and_sell_stock"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    Number
	expected int
}

type Number struct {
	prices []int
}

func TestMyAtoi(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			expected: 5,
		},
		{
			name: "test2",
			input: Number{
				prices: []int{7, 6, 4, 3, 1},
			},
			expected: 0,
		},
		{
			name: "test3",
			input: Number{
				prices: []int{4, 3, 2, 6, 4},
			},
			expected: 4,
		},
		{
			name: "test4",
			input: Number{
				prices: []int{2, 4, 1},
			},
			expected: 2,
		},
		{
			name: "test5",
			input: Number{
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			expected: 4,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MaxProfit(tt.input.prices))
		})
	}
}
