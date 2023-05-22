package float_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/float"
	"github.com/stretchr/testify/assert"
)

func TestPantul(t *testing.T) {
	t.Run("Test Pantul", func(t *testing.T) {
		assert.Equal(t, "0021350", Solution(7, 213.5))
		assert.Equal(t, "00100", Solution(5, 1))

	})
}
