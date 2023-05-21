package faktorial_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/faktorial"
	"github.com/stretchr/testify/assert"
)

func TestFaktorial(t *testing.T) {
	t.Run("test faktorial", func(t *testing.T) {
		assert.Equal(t, 120, Faktorial(5))
		assert.Equal(t, 720, Faktorial(6))
	})
}
