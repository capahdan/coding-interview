package array_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/array"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, Combine([]int{1, 2, 3}, []int{4, 5, 6}))
	})
}
