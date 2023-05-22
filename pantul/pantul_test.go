package pantul_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/pantul"
	"github.com/stretchr/testify/assert"
)

func TestPantul(t *testing.T) {
	t.Run("Test Pantul", func(t *testing.T) {
		assert.Equal(t, []string{"4", "2", "1", "0.5"}, Pantul(4))
		assert.Equal(t, []string{"7", "2", "1", "0.5"}, Pantul(4))
	})
}
