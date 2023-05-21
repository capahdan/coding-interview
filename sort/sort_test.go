package sort_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/sort"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Run("coba sorting", func(t *testing.T) {
		assert.Equal(t, []int{9, 8, 5, 2, 1}, Sort([]int{5, 2, 8, 1, 9}))
	})
}
