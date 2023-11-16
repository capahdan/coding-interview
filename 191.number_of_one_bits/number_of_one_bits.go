package number_of_one_bits

import "fmt"

func HammingWeight(num uint32) int {
	// bitStr := fmt.Sprintf("%032b", num)

	// oneCount := 0

	// for i := range bitStr {
	// 	if bitStr[i] == '1' {
	// 		oneCount++
	// 	}
	// }

	// return oneCount

	count := 0
	for i := 0; i < 32; i++ {
		// Gunakan bitwise AND untuk mengecek apakah bit ke-i adalah 1

		fmt.Println(1 << i)

		fmt.Println(num & (1 << i))
		if num&(1<<i) != 0 {
			count++
		}
	}
	return count
}

// ganti dia menjadi binary
// hitung angka 1 ada berapa banyak

// maximal angka dari 32 bit adalah

// 4294967295
