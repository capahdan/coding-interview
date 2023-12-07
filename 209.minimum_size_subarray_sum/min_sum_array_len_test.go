package min_sum_array_len_test

import (
	"testing"

	. "github.com/capahdan/coding-interview/209.minimum_size_subarray_sum"
	"github.com/stretchr/testify/assert"
)

func TestMinSumArrayLenTest(t *testing.T) {

}

type TestTable struct {
	name     string
	input    Number
	expected int
}

type Number struct {
	target int
	nums   []int
}

func TestMinSubArrayLen(t *testing.T) {
	testTable := []TestTable{
		{
			name: "test1",
			input: Number{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			expected: 2,
		},
		{
			name: "test2",
			input: Number{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			expected: 1,
		},
		{
			name: "test3",
			input: Number{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			expected: 0,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MinSubArrayLen(tt.input.target, tt.input.nums))
		})
	}
}

// func MinSubArrayLen(target int, nums []int) int {

// 	minLength := math.MaxInt32
// 	sum, start := 0, 0

// 	for end := 0; end < len(nums); end++ {
// 		sum += nums[end]

// 		for sum >= target {
// 			if end-start+1 < minLength {
// 				minLength = end - start + 1
// 			}
// 			sum -= nums[start]
// 			start++
// 		}
// 	}

// 	if minLength == math.MaxInt32 {
// 		return 0
// 	}
// 	return minLength
// }
