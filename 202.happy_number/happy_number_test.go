package happy_number_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/202.happy_number"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    int
	expected bool
}

func TestHappyNumber(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    19,
			expected: true,
		},
		// {
		// 	name:     "test2",
		// 	input:    2,
		// 	expected: false,
		// },
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsHappy(tt.input))
		})
	}
}
