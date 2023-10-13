package happy_number

import "fmt"

// first solution
// func IsHappy(n int) bool {

// 	seen := make(map[int]bool)
// 	for n != 1 && !seen[n] {
// 		sum := 0

// 		seen[n] = true

// 		for n != 0 {
// 			remainder := n % 10
// 			sum += remainder * remainder
// 			n /= 10
// 		}

// 		n = sum

// 	}

// 	return n == 1
// }

// second solution

func nextNumber(n int) int {
	var newNumber int
	for n != 0 {
		remainder := n % 10
		newNumber += remainder * remainder
		n /= 10
	}
	return newNumber
}

func IsHappy(n int) bool {
	slowPointer := n
	fastPointer := nextNumber(n)
	loop := 1

	for fastPointer != 1 && fastPointer != slowPointer {
		slowPointer = nextNumber(slowPointer)
		fastPointer = nextNumber(nextNumber(fastPointer))
		fmt.Printf("\nfirst loop: %d , slowPointer:%d , fastPointer:%d", loop, slowPointer, fastPointer)

		loop++
	}
	return fastPointer == 1
}
