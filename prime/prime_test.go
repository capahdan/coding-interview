package prime_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/prime"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19}, Prime(2, 20))
		assert.Equal(t, []int{2, 3, 5}, Prime(2, 6))
	})
}
