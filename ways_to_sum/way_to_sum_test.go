package way_to_sum_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/ways_to_sum"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    num
	expected int
}

type num struct {
	k     int
	total int
}

func TestSolu(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test 1",
			input: num{
				k:     2,
				total: 8,
			},
			expected: 5,
		},
		// {
		// 	name: "test 2",
		// 	input: num{
		// 		k:     2,
		// 		total: 6,
		// 	},

		// 	expected: 4,
		// },

		// {
		// 	name: "test 3",
		// 	input: num{
		// 		k:     3,
		// 		total: 7,
		// 	},

		// 	expected: 8,
		// },
		// {
		// 	name: "test 4",
		// 	input: num{
		// 		k:     2,
		// 		total: 2,
		// 	},

		// 	expected: 2,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, WayToSum(tt.input.total, tt.input.k))
		})
	}
}
