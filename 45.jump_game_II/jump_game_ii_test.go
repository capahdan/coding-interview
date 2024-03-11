package jump_game_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/45.jump_game_II"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   []int
	output  int
}

func TestJumpGame(t *testing.T) {

	testtable := []testTable{

		{
			message: "test1",
			input:   []int{2, 3, 1, 1, 4},
			output:  2,
		},
		// {
		// 	message: "test 2",
		// 	input:   []int{2, 3, 0, 1, 4},
		// 	output:  2,
		// },
	}

	for _, tt := range testtable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, JumpGame(tt.input))
		})
	}

}
