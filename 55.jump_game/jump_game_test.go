package jump_game_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/55.jump_game"
	"github.com/stretchr/testify/assert"
)

type TestTable struct {
	name     string
	input    []int
	expected bool
}

func TestJumpGame(t *testing.T) {
	testTable := []TestTable{
		{
			name:     "test1",
			input:    []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "test2",
			input:    []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CanJump(tt.input))
		})
	}
}
