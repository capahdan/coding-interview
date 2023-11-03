package add_binary

import "strings"

func AddBinary(a string, b string) string {
	var res strings.Builder
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		res.WriteByte(byte(sum%2) + '0')
	}
	if carry != 0 {
		res.WriteByte(byte(carry) + '0')
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
