package two_sum_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/two_sum"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		assert.Equal(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	})
}
