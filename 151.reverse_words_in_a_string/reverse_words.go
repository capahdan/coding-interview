package reverse_word

import "strings"

func ReverseWord(s string) string {
	// arrString := strings.Split(s, " ")

	// var result string

	// for i := len(arrString) - 1; i >= 0; i-- {
	// 	if arrString[i] != "" {
	// 		result += " " + arrString[i]
	// 	}
	// }

	// result = strings.TrimLeft(result, " ")

	// return result

	// chat gpt solution

	// Step 1: Trim leading and trailing spaces
	trimmedStr := strings.TrimSpace(s)

	// Step 2: Split the string into words
	words := strings.Fields(trimmedStr)

	// Step 3: Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Step 4: Join the words back together
	return strings.Join(words, " ")
}

//  yang perlu kita lakukan adalah menghapus semua spasi baik depan maupun dari belakang
//  lalu kita split saja dulu stringnya menggunakan spasi
// lalu kita buat string baru dengan menambahkan spasi juga dari belakang ke depan

// solusi kedua
// kita buat stack of string
// untuk mengisi stack itu kita for loop dulu kalo dia menemukan
