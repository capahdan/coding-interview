package sqrt_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/69.sqrt"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    int
	expected int
}

func TestMySqrt(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test 1",
			input:    4,
			expected: 2,
		},
		{
			name:     "test 2",
			input:    8,
			expected: 2,
		},
		{
			name:     "test 3",
			input:    9,
			expected: 3,
		},
		{
			name:     "test 4",
			input:    12,
			expected: 3,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MySqrt(tt.input))
		})
	}
}
