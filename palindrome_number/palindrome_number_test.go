package palindrome_number_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/palindrome_number"
	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumber(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		assert.True(t, IsPalindrome(0))
	})
	t.Run("negative number", func(t *testing.T) {
		assert.False(t, IsPalindrome(-1))
	})
	t.Run("positive number", func(t *testing.T) {
		assert.True(t, IsPalindrome(121))
	})
	t.Run("wrong Number", func(t *testing.T) {
		assert.False(t, IsPalindrome(123))
	})
}
