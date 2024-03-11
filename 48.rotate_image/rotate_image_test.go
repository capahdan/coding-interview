package rotate_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/48.rotate_image"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   [][]int
	output  [][]int
}

func TestRotateMatriks(t *testing.T) {

	testtable := []testTable{

		{
			message: "test 3 * 3 matriks ",
			input:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output:  [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		// {
		// 	message: "test 4 * 4 matriks ",
		// 	input:   [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		// 	output:  [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		// },
	}

	for _, tt := range testtable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, RotateImage(tt.input))
		})
	}

}
