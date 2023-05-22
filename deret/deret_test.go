package deret_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/deret"
	"github.com/stretchr/testify/assert"
)

func TestPantul(t *testing.T) {
	t.Run("Test Pantul", func(t *testing.T) {
		assert.Equal(t, []int{-1, 2, 1, 3, 4, 7}, Deret(-1, 2, 6))
		assert.Equal(t, []int{-1, 2, 1, 3, 4, 7}, Deret(13, 18, 10))
	})
}
