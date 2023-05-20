package palindrome_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/palindrome"
	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		assert.True(t, Palindrome(""))
	})
	t.Run("single character", func(t *testing.T) {
		assert.True(t, Palindrome("a"))
	})
	t.Run("racecar", func(t *testing.T) {
		assert.True(t, Palindrome("racecar"))
	})
}
