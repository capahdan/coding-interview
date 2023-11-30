package integer_to_roman

func IntToRoman(num int) string {
	// var result string

	// arrInt := []int{1000, 500, 100, 50, 10, 5, 1}

	// for num >= 1 {
	// 	// kita kurang numnya
	// 	for _, value := range arrInt {
	// 		temp := num - value
	// 		if temp >= 0 {
	// 			num -= value
	// 			result += romanMap[value]
	// 			// arrInt = arrInt[idx:]
	// 		}
	// 	}
	// }

	// return result
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

// yang pasti kita harus punya map yang berisikan nilai dari setiap karaketer

var romanMap = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

// yang akan saya lakukan adalah
// saya akan mencoba untuk melooping nilai dari input sampai hasilnya akan menjadi 0

// 3
// for num <1
// looping for key, value := romanMap
