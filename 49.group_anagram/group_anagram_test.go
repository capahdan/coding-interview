package group_anagram_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/49.group_anagram"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	message string
	input   []string
	output  [][]string
}

func TestIntegerToRoman(t *testing.T) {
	testTable := []testTable{
		{
			message: "test1",
			input:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			message: "test1",
			input:   []string{""},
			output: [][]string{
				{""},
			},
		},
		{
			message: "test1",
			input:   []string{"a"},
			output: [][]string{
				{"a"},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.message, func(t *testing.T) {
			assert.Equal(t, tt.output, GroupAnagrams(tt.input))
		})
	}
}
