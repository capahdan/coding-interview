package word_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/word"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		assert.Equal(t, 6, WordCount("Hari ini adalah hari yang cerah"))
		assert.Equal(t, 1, WordCount("Hari"))
		assert.Equal(t, 2, WordCount("Hari ini"))
	})
}
