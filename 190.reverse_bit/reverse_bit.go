package reverse_bit

import "strings"

// func ReverseBit(num uint32) uint32 {

// 	binaryStr := parseToBit(num)
// 	paddedStr := fmt.Sprintf("%032s", binaryStr)

// 	reversedStr := reverse(paddedStr)

// 	reversedInt := strToInt(reversedStr)

// 	return reversedInt
// }

func parseToBit(num int) string {
	var res strings.Builder

	for num != 0 {
		num = num / 2
		res.WriteByte(byte(num%2) + '0')
	}

	result := res.String()
	return reverse(result)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// func strToInt(str string) uint32 {
// 	var result uint32

// 	val, err := strconv.ParseUint(str, 10, 32)
// 	if err != nil {
// 		// Handle error sesuai kebutuhan, di sini kita hanya mencetak pesan error dan mengembalikan 0
// 		fmt.Println("Error converting string to uint32:", err)
// 		return 0
// 	}
// 	// Karena kita tahu bahwa nilai yang dikembalikan oleh ParseUint sesuai dengan ukuran uint32,
// 	// kita bisa langsung melakukan konversi ke tipe tersebut.
// 	result = uint32(val)
// 	return result
// }

func ReverseBit(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
