package ganjil_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/ganjil"
	"github.com/stretchr/testify/assert"
)

func TestGanjil(t *testing.T) {
	t.Run("test ganjil", func(t *testing.T) {
		assert.Equal(t, 5, Ganjil([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		assert.Equal(t, 3, Ganjil([]int{1, 3, 9, 10}))
	})
}
