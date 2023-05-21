package main

import (
	"fmt"

	. "github.com/capahdan/coding-interview/faktorial"
	. "github.com/capahdan/coding-interview/sort"
	. "github.com/capahdan/coding-interview/two_sum"
)

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(Faktorial(5))
	fmt.Println(Sort([]int{5, 2, 8, 1, 9}))
}
